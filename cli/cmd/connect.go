package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to gRPC service",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}
		defer conn.Close()

		client := (conn)
		response, err := client.YourMethod(context.Background(), &YourRequest{})
		if err != nil {
			log.Fatalf("Error while calling gRPC: %v", err)
		}
		fmt.Println("Response from server:", response)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}
