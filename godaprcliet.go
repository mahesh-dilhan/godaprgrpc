package main

import (
	"context"
	"log"
	"os"

	dapr "github.com/dapr/go-sdk/client"
)

func main() {

	// just for this demo
	ctx := context.Background()
	data := []byte("ping")

	// create the client
	client, err := dapr.NewClient()
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	// save state with the key key1
	err = client.SaveState(ctx, "statestore", "key1", data)
	if err != nil {
		log.Panic(err)
	}
	log.Println("data saved")
}
