package main

import (
	"context"
	"log"
	"net"

	common "github.com/piyushbag/oms-common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer()
	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	defer l.Close()

	store := NewStore()
	service := NewService(store)
	NewGRPCHandler(grpcServer)

	service.CreateOrder(context.Background())

	log.Println("Starting GRPC server at", grpcAddr)

	// register service
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal("Failed to serve: ", err)
	}
}
