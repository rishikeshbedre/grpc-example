package main

import (
	"grpc-example/client"
	"grpc-example/server"
	"time"
)

func main() {
	go server.StartServer()

	time.Sleep(5 * time.Second)

	client.StartClient()
}
