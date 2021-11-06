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

func eventHandler(ctx context.Context, e *common.TopicEvent) (retry bool, err error) {
	log.Printf("event - PubsubName:%s, Topic:%s, ID:%s, Data: %v", e.PubsubName, e.Topic, e.ID, e.Data)
	// do something with the event
	return true, nil
}

func echoHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	log.Printf("echo - ContentType:%s, Verb:%s, QueryString:%s, %+v", in.ContentType, in.Verb, in.QueryString, string(in.Data))
	// do something with the invocation here
	out = &common.Content{
		Data:        in.Data,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}
