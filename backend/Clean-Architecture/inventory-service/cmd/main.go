package main

import (
	"inventory-service/db"
	"inventory-service/internal/inventory/cache"
	"inventory-service/internal/product"

	"log"
	"net"

	pb "github.com/elliaaan/proto-gen/pb/inventory/github.com/elliaaan/proto-gen/pb/inventory"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	database := db.InitDB()

	itemCache := cache.NewCache(database)

	repo := &product.Repository{DB: database}
	service := &product.Service{
		Repo:  repo,
		Cache: itemCache,
	}

	server := &product.GRPCServer{Service: service}

	grpcServer := grpc.NewServer()
	pb.RegisterInventoryServiceServer(grpcServer, server)

	log.Println("Inventory gRPC service running on port 8080...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
