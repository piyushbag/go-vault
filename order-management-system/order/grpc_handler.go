package main

import (
	"context"
	"log"

	pb "github.com/piyushbag/oms-common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New order received, Order: %v", p)

	// create order
	o := &pb.Order{
		OrderId: "42",
	}
	return o, nil
}
