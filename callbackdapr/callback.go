package main

import (
	daprd "github.com/dapr/go-sdk/service/grpc"
	"log"
	"net"
)

func main() {
	list, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatalf("gRPC listener creation failed: %s", err)
	}
	s := daprd.NewServiceWithListener(list)

	if err := s.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}

	sub := &common.Subscription{
		PubsubName: "messages",
		Topic:      "topic1",
	}
	if err := s.AddTopicEventHandler(sub, eventHandler); err != nil {
		log.Fatalf("error adding topic subscription: %v", err)
	}
}
