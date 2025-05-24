package main

import (
	"log"
	"net"

	"order-service/db"
	"order-service/internal/order"

	pb "github.com/elliaaan/proto-gen/pb/order/github.com/elliaaan/proto-gen/pb/order"
	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {

	database := db.InitDB()

	natsConn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS: %v", err)
	}

	orderCache := order.NewCache(database)

	repo := &order.Repository{DB: database}
	service := &order.Service{
		Repo:  repo,
		Cache: orderCache,
		NATS:  natsConn,
	}
	server := &order.GRPCServer{Service: service}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Order Service running on port 8081")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
