package main

import (
	"fmt"
	"log"

	"github.com/go-ole/go-ole"
)

func main() {
	// Инициализация OLE
	err := ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	if err != nil {
		log.Fatalf("Failed to initialize OLE: %v\n", err)
	}
	defer ole.CoUninitialize()

	// Запуск мониторинга создания процессов
	go func() {
		processCreatedCh, processErrorCh := monitorProcessCreation()
		for {
			select {
			case process := <-processCreatedCh:
				fmt.Printf("[Process Created] %+v\n", process)
			case err := <-processErrorCh:
				if err != nil {
					fmt.Printf("[Error - Process Created] %v\n", err)
				}
			}
		}
	}()

	// Запуск мониторинга удаления процессов
	go func() {
		processDeletedCh, processDelErrorCh := monitorProcessDeletion()
		for {
			select {
			case process := <-processDeletedCh:
				fmt.Printf("[Process Deleted] %+v\n", process)
			case err := <-processDelErrorCh:
				if err != nil {
					fmt.Printf("[Error - Process Deleted] %v\n", err)
				}
			}
		}
	}()

	//Запуск мониторинга evtx
	go func() {
		logCh, logErrCh := monitorEvtx()
		for {
			select {
			case logEvent := <-logCh:
				fmt.Printf("[Log Event - %s]\n", logEvent.Message)
			case err := <-logErrCh:
				if err != nil {
					fmt.Printf("[Error - Log Event - %s]\n", err)
				}
			}
		}
	}()

	select {}
}
