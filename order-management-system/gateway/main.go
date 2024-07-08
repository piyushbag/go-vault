package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	common "github.com/piyushbag/oms-common"
	pb "github.com/piyushbag/oms-common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr     = common.EnvString("HTTP_ADDR", ":8080")
	orderSvcAddr = "localhost:2000"
)

func main() {
	conn, err := grpc.NewClient(orderSvcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to dial order service: ", err)
	}
	defer conn.Close()

	log.Println("Dialing order service at", orderSvcAddr)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Println("Starting server on", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to listen and serve: ", err)
	}
}
