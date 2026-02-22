package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ably/ably-go/ably"
)

func main() {
	client, err := ably.NewRealtime(
		ably.WithKey("toGIsA.95eGfg:mos7cWjFjKSeWk2tEDFmLRHHe3mkm5Uh2IJIX2MMPns"),
		ably.WithClientID("my-first-client"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Wait for the connection to be connected
	connStateChan := make(chan ably.ConnectionStateChange, 1)
	client.Connection.On(ably.ConnectionEventConnected, func(change ably.ConnectionStateChange) {
		connStateChan <- change
	})
	select {
	case <-connStateChan:
		fmt.Println("Made my first connection!")
	case <-context.Background().Done():
		log.Fatal("Context cancelled before connection established")
	}

	// Keep the program running
	select {}
}
