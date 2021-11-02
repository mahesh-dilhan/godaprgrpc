package main

import (
	daprd "github.com/dapr/go-sdk/service/grpc"
	"log"
)

func main() {
	s, err := daprd.NewService(":50001")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}
