package main

import (
	"fmt"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// Структура для представления процесса
type Win32Process struct {
	ProcessID   int32
	Name        string
	CommandLine string
}

// Структура для представления события журнала
type Win32NTLogEvent struct {
	EventID  int32
	Logfile  string
	Message  string
	Category string
	Source   string
}

// Функция для мониторинга создания процессов
func monitorProcessCreation() (<-chan Win32Process, <-chan error) {
	return monitorProcesses("__InstanceCreationEvent")
}

// Функция для мониторинга удаления процессов
func monitorProcessDeletion() (<-chan Win32Process, <-chan error) {
	return monitorProcesses("__InstanceDeletionEvent")
}

// Универсальная функция для мониторинга процессов
func monitorProcesses(eventClass string) (<-chan Win32Process, <-chan error) {
	processCh := make(chan Win32Process)
	errorCh := make(chan error)

	go func() {
		defer close(processCh)
		defer close(errorCh)

		// Подключение к WMI через SWbemLocator
		unknown, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
		if err != nil {
			errorCh <- fmt.Errorf("Failed to create SWbemLocator object: %v", err)
			return
		}
		defer unknown.Release()

		wmi, err := unknown.QueryInterface(ole.IID_IDispatch)
		if err != nil {
			errorCh <- fmt.Errorf("Failed to query SWbemLocator interface: %v", err)
			return
		}
		defer wmi.Release()

		serviceRaw, err := oleutil.CallMethod(wmi, "ConnectServer", nil, "root\\cimv2")
		if err != nil {
			errorCh <- fmt.Errorf("Failed to connect to WMI service: %v", err)
			return
		}
		service := serviceRaw.ToIDispatch()
		defer service.Release()

		// Подписка на события процессов
		query := fmt.Sprintf(`
			SELECT * 
			FROM %s 
			WITHIN 1 
			WHERE TargetInstance ISA 'Win32_Process'`, eventClass)

		eventSinkRaw, err := oleutil.CallMethod(service, "ExecNotificationQuery", query)
		if err != nil {
			errorCh <- fmt.Errorf("Failed to execute WMI notification query: %v", err)
			return
		}
		eventSink := eventSinkRaw.ToIDispatch()
		defer eventSink.Release()

		// Постоянный цикл для обработки событий процессов
		for {
			eventRaw, err := oleutil.CallMethod(eventSink, "NextEvent", 0xFFFFFFFF)
			if err != nil {
				errorCh <- fmt.Errorf("Error receiving process event: %v", err)
				continue
			}
			event := eventRaw.ToIDispatch()
			defer event.Release()

			targetInstanceRaw, err := oleutil.GetProperty(event, "TargetInstance")
			if err != nil {
				errorCh <- fmt.Errorf("Failed to get TargetInstance property: %v", err)
				continue
			}
			targetInstance := targetInstanceRaw.ToIDispatch()
			defer targetInstance.Release()

			process := Win32Process{
				ProcessID:   int32(oleutil.MustGetProperty(targetInstance, "ProcessId").Val),
				Name:        oleutil.MustGetProperty(targetInstance, "Name").ToString(),
				CommandLine: oleutil.MustGetProperty(targetInstance, "CommandLine").ToString(),
			}

			processCh <- process
		}
	}()

	return processCh, errorCh
}

// Функция для мониторинга Evtx
func monitorEvtx() (<-chan Win32NTLogEvent, <-chan error) {
	logCh := make(chan Win32NTLogEvent)
	errorCh := make(chan error)

	go func() {
		defer close(logCh)
		defer close(errorCh)

		// Подключение к WMI через SWbemLocator
		unknown, err := oleutil.CreateObject("WbemScripting.SWbemLocator")
		if err != nil {
			errorCh <- fmt.Errorf("Failed to create SWbemLocator object for log files: %v", err)
			return
		}
		defer unknown.Release()

		wmi, err := unknown.QueryInterface(ole.IID_IDispatch)
		if err != nil {
			errorCh <- fmt.Errorf("Failed to query SWbemLocator interface for log files: %v", err)
			return
		}
		defer wmi.Release()

		serviceRaw, err := oleutil.CallMethod(wmi, "ConnectServer", nil, "root\\cimv2")
		if err != nil {
			errorCh <- fmt.Errorf("Failed to connect to WMI service for log files: %v", err)
			return
		}
		service := serviceRaw.ToIDispatch()
		defer service.Release()

		// Формирование запроса для одного лог-файла
		query := fmt.Sprintf(`
			SELECT * 
			FROM __InstanceCreationEvent 
			WITHIN 1 
			WHERE TargetInstance ISA 'Win32_NTLogEvent'`)

		eventSinkRaw, err := oleutil.CallMethod(service, "ExecNotificationQuery", query)
		if err != nil {
			errorCh <- fmt.Errorf("Failed to execute WMI query for log files: %v", err)
			return
		}
		eventSink := eventSinkRaw.ToIDispatch()
		defer eventSink.Release()

		// Постоянный цикл для обработки событий журнала
		for {
			eventRaw, err := oleutil.CallMethod(eventSink, "NextEvent", 0xFFFFFFFF)
			if err != nil {
				errorCh <- fmt.Errorf("Error receiving log files: %v", err)
				continue
			}
			event := eventRaw.ToIDispatch()
			defer event.Release()

			targetInstanceRaw, err := oleutil.GetProperty(event, "TargetInstance")
			if err != nil {
				errorCh <- fmt.Errorf("Failed to get TargetInstance property for log files%v", err)
				continue
			}
			targetInstance := targetInstanceRaw.ToIDispatch()
			defer targetInstance.Release()

			logEvent := Win32NTLogEvent{
				EventID:  int32(oleutil.MustGetProperty(targetInstance, "EventCode").Val),
				Logfile:  safeGetStringProperty(targetInstance, "Logfile"),
				Message:  safeGetStringProperty(targetInstance, "Message"),
				Category: safeGetStringProperty(targetInstance, "CategoryString"),
				Source:   safeGetStringProperty(targetInstance, "SourceName"),
			}

			logCh <- logEvent
		}
	}()

	return logCh, errorCh
}

// Функция для безопасного извлечения строкового свойства
func safeGetStringProperty(obj *ole.IDispatch, propName string) string {
	prop, err := oleutil.GetProperty(obj, propName)
	if err != nil || prop.VT == ole.VT_NULL || prop.VT == ole.VT_EMPTY {
		return "" // Возвращаем пустую строку, если свойство не существует
	}
	return prop.ToString()
}
