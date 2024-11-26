package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-ole/go-ole"
	pb "github.com/imyazip/GoSIEM/windows_agent/proto"
	"google.golang.org/grpc"
)

func main() {
	/*
		apiKey := "your-api-key"
		token, err := getToken(apiKey)
		if err != nil {
			log.Fatalf("Failed to get token: %v", err)
		}

		// Выводим полученный токен
		fmt.Printf("Received token: %s\n", token)
		// Инициализация OLE
	*/

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
				//fmt.Printf("[Log Event - %s]\n", logEvent.InsertionStrings)
				fmt.Println(logEvent.EventID)
				fmt.Println(logEvent.Message)
				for i, value := range logEvent.InsertionStrings {
					fmt.Println(i, ": [", value, "]")
				}
			case err := <-logErrCh:
				if err != nil {
					fmt.Printf("[Error - Log Event - %s]\n", err)
				}
			}
		}
	}()

	select {}
}

// Функция для получения токена
func getToken(apiKey string) (string, error) {
	// Создаем соединение с gRPC сервером
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // Замените на актуальный адрес
	if err != nil {
		return "", fmt.Errorf("did not connect: %v", err)
	}
	defer conn.Close()

	// Создаем клиент для сервиса
	client := pb.NewAuthServiceClient(conn)

	// Создаем запрос с API-ключом
	req := &pb.GenerateJWTForSensorRequest{
		ApiKey: apiKey,
	}

	// Отправляем запрос и получаем ответ
	res, err := client.GenerateJWTForSensor(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("error during request: %v", err)
	}

	return res.Token, nil
}
