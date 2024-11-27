package parser

import (
	"fmt"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// Структура для представления процесса
type Win32_Process struct {
	ProcessID   int32
	Name        string
	CommandLine string
}

// Структура для представления события журнала
type Win32_NTLogEvent struct {
	EventID          int32
	Logfile          string
	Message          string
	Category         string
	Source           string
	InsertionStrings []string
}

// Функция для мониторинга создания процессов
func MonitorProcessCreation() (<-chan Win32_Process, <-chan error) {
	return monitorProcesses("__InstanceCreationEvent")
}

// Функция для мониторинга удаления процессов
func MonitorProcessDeletion() (<-chan Win32_Process, <-chan error) {
	return monitorProcesses("__InstanceDeletionEvent")
}

// Универсальная функция для мониторинга процессов
func monitorProcesses(eventClass string) (<-chan Win32_Process, <-chan error) {
	processCh := make(chan Win32_Process)
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

			process := Win32_Process{
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
func MonitorEvtx() (<-chan Win32_NTLogEvent, <-chan error) {
	logCh := make(chan Win32_NTLogEvent)
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

			logEvent := Win32_NTLogEvent{
				EventID:          int32(oleutil.MustGetProperty(targetInstance, "EventCode").Val),
				Logfile:          safeGetStringProperty(targetInstance, "Logfile"),
				Message:          safeGetStringProperty(targetInstance, "Message"),
				Category:         safeGetStringProperty(targetInstance, "CategoryString"),
				Source:           safeGetStringProperty(targetInstance, "SourceName"),
				InsertionStrings: extractInsertionStrings(targetInstance),
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

// Функция для извлечения InsertionStrings из объекта WMI
func extractInsertionStrings(targetInstance *ole.IDispatch) []string {
	insertionStrings := []string{}

	// Попытка получить свойство InsertionStrings
	insertionStringsRaw, err := oleutil.GetProperty(targetInstance, "InsertionStrings")
	if err != nil {
		fmt.Printf("Error fetching InsertionStrings property: %v\n", err)
		return insertionStrings
	}

	// Проверка типа данных
	if insertionStringsRaw.VT != (ole.VT_ARRAY | ole.VT_VARIANT) {
		fmt.Printf("Unexpected variant type: VT(%d)\n", insertionStringsRaw.VT)
		return insertionStrings
	}

	// Преобразование в SafeArray
	safeArray := insertionStringsRaw.ToArray()
	if safeArray == nil {
		fmt.Println("SafeArray conversion failed.")
		return insertionStrings
	}
	defer safeArray.Release()

	// Преобразование SafeArray в массив значений
	values := safeArray.ToValueArray()
	if len(values) == 0 {
		fmt.Println("No values in SafeArray.")
		return insertionStrings
	}

	// Обработка каждого элемента массива
	for _, value := range values {
		// Проверяем, что элемент является строкой
		switch v := value.(type) {
		case string:
			insertionStrings = append(insertionStrings, v)
		case *ole.VARIANT:
			if v.VT == ole.VT_BSTR {
				insertionStrings = append(insertionStrings, v.ToString())
			} else {
				fmt.Printf("Unsupported variant type in array: VT(%d)\n", v.VT)
			}
		default:
			fmt.Printf("Non-string value found in SafeArray: %v (type %T)\n", value, value)
		}
	}

	return insertionStrings
}
