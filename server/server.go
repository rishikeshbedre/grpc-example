package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"grpc-example/pb/employeePB"
)

type exampleServer struct {
	employeePB.UnimplementedEmployeeServiceServer
}

func (s *exampleServer) SayHello(ctx context.Context, input *employeePB.HelloRequest) (*employeePB.HelloReply, error) {
	log.Println("Request name: ", input.Name)
	tempReply := &employeePB.HelloReply{
		Message: "Hello " + input.GetName(),
	}
	return tempReply, nil
}

func StartServer() {
	lis, listenErr := net.Listen("tcp", ":5000")
	if listenErr != nil {
		println(listenErr)
	}

	grpcServer := grpc.NewServer()
	employeePB.RegisterEmployeeServiceServer(grpcServer, &exampleServer{})
	log.Println("Server listening at ", lis.Addr())

	grpcServeErr := grpcServer.Serve(lis)
	if grpcServeErr != nil {
		println(grpcServeErr)
	}
}
