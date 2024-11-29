package main

import (
	"fmt"
	"log"

	"github.com/go-ole/go-ole"
	"github.com/imyazip/GoSIEM/windows_agent/internal/pkg/api"
	"github.com/imyazip/GoSIEM/windows_agent/internal/pkg/config"
	"github.com/imyazip/GoSIEM/windows_agent/internal/pkg/parser"
	"github.com/imyazip/GoSIEM/windows_agent/internal/pkg/utils"
)

func main() {
	cfg := config.LoadConfig("agent-config.yaml")
	conn, err := api.ConnectToServer(fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		log.Fatal(err.Error())
	}
	authClient := api.GetAuthClient(conn)
	log.Printf("Connected to AuthService")
	logClient := api.GetLogClient(conn)
	log.Printf("Connected to LogService")

	token, err := api.GetToken(cfg.Api.Key, authClient)
	if err != nil {
		log.Fatalf("Failed to get token: %v", err)
	}

	ctx := api.CreateAuthContext(token)

	// Выводим полученный токен
	fmt.Printf("Received token: %s\n", token)
	// Инициализация OLE

	err = ole.CoInitializeEx(0, ole.COINIT_MULTITHREADED)
	if err != nil {
		log.Fatalf("Failed to initialize OLE: %v\n", err)
	}
	defer ole.CoUninitialize()

	/* Запуск мониторинга создания процессов
	go func() {
		processCreatedCh, processErrorCh := parser.MonitorProcessCreation()
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
		processDeletedCh, processDelErrorCh := parser.MonitorProcessDeletion()
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
	*/
	//Запуск мониторинга evtx
	go func() {
		logCh, logErrCh := parser.MonitorEvtx()
		for {
			select {
			case logEvent := <-logCh:
				serialized := utils.StructToSlice(logEvent)
				api.SendSerializedLog(ctx, serialized, "sensor", "evtx", logClient)
			case err := <-logErrCh:
				if err != nil {
					fmt.Printf("[Error - Log Event - %s]\n", err)
				}
			}
		}
	}()

	select {}
}
