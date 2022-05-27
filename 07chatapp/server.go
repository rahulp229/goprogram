package main

import (
	"07chatapp/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen port 9000: %v", err)
	}
	s := chat.Server{}
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, &s)
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc on port 9000: %v", err)
	}

}
