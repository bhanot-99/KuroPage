package main

import (
	"encoding/json"
	"log"
	"net"
	"time"

	"github.com/elliaaan/statistics-service/db"
	"github.com/elliaaan/statistics-service/internal/statistics"
	"github.com/elliaaan/statistics-service/models"

	statisticspb "github.com/elliaaan/proto-gen/pb/statistics"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	database := db.InitDB()
	if err := database.AutoMigrate(&models.Event{}); err != nil {
		log.Fatalf("Ошибка миграции таблицы Event: %v", err)
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS: %v", err)
	}
	defer nc.Drain()

	statistics.SubscribeToOrderCreated(nc, database)

	service := &statistics.Service{DB: database}
	server := &statistics.GRPCServer{Service: service}

	grpcServer := grpc.NewServer()
	statisticspb.RegisterStatisticsServiceServer(grpcServer, server)

	startPeriodicPublisher(nc)

	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Ошибка запуска gRPC-сервера: %v", err)
	}

	log.Println(" Statistics Service запущен на порту 8082")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка gRPC Serve: %v", err)
	}
}

func startPeriodicPublisher(nc *nats.Conn) {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for t := range ticker.C {
			message := map[string]interface{}{
				"source": "statistics-service",
				"type":   "update",
				"time":   t.Format(time.RFC3339),
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Printf("Ошибка сериализации сообщения: %v", err)
				continue
			}

			err = nc.Publish("ap2.statistics.event.updated", data)
			if err != nil {
				log.Printf("Ошибка публикации в NATS: %v", err)
			} else {
				log.Println(" Отправлено сообщение в ap2.statistics.event.updated:", string(data))
			}
		}
	}()
}
