package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "5000"
	}

	listner, err := net.Listen("tcp", ":"+Port)
	if err != nil {
		log.Fatalf("cound not listen at %v and error : %v ", Port, err)
	}

	log.Println("listening port : ", Port)

	//gRPC server instance

	grpcServer := grpc.NewServer()

	//gRPC listen and serve
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatalf("failed to start gRPC server : %v ", err)
	}
}
