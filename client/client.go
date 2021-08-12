package client

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"grpc-example/pb/employeePB"
)

func StartClient() {
	conn, connErr := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if connErr != nil {
		log.Println(connErr)
	}
	defer conn.Close()

	grpcClient := employeePB.NewEmployeeServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &employeePB.HelloRequest{
		Name: "Rishi",
	}
	resp, reqErr := grpcClient.SayHello(ctx, req)
	if reqErr != nil {
		log.Println(reqErr)
	}
	log.Println("Message received: ", resp.Message)
}
