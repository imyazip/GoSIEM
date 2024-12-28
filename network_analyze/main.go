package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/imyazip/network_analyze/config"
	"github.com/imyazip/sigolyze"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/imyazip/network_analyze/proto"
)

func main() {
	cfg := config.LoadConfig("analyzer.yaml")
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.LogServer.Host, cfg.LogServer.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewLogStorageServiceClient(conn)

	compiler := sigolyze.NewCompiler()
	compiler.LoadSignatureFromJson("./signatures/sig.json")

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <interface>")
		os.Exit(1)
	}

	// Получение интерфейса из аргументов командной строки
	interfaceName := os.Args[1]

	// Открытие устройства для прослушивания
	handle, err := pcap.OpenLive(interfaceName, 65536, true, pcap.BlockForever)
	if err != nil {
		log.Fatalf("Error opening device %s: %v", interfaceName, err)
	}
	defer handle.Close()

	// Установка фильтра для IP-пакетов (опционально)
	err = handle.SetBPFFilter("ip")
	if err != nil {
		log.Fatalf("Error setting BPF filter: %v", err)
	}

	fmt.Printf("Listening on interface %s for IP packets...\n", interfaceName)

	// Создание канала пакетов
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		matches := sigolyze.Match(compiler, packet.String())
		if len(matches) != 0 {
			log.Printf("Matched with signature: %s", matches[0].Name)
			req := &pb.AddSecurityEventRequest{
				LogId:            0,
				EventType:        matches[0].Name,
				EventDescription: "",
			}

			_, err := pb.LogStorageServiceClient.AddSecurityEvent(client, context.Background(), req)
			if err != nil {
				log.Printf("Error sending event: %s", err)
			}
		}
	}

}
