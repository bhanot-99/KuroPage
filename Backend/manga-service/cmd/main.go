package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/bhanot-99/KuroPage/Backend/manga-service/internal/service"
	"github.com/bhanot-99/KuroPage/Backend/pkg/config"
	"github.com/bhanot-99/KuroPage/Backend/pkg/logger"
	"github.com/bhanot-99/KuroPage/Backend/pkg/natsutils"
	"github.com/bhanot-99/KuroPage/Backend/pkg/proto"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	logger.Init()
	defer logger.Sync()

	cfg := config.Load()

	// Database connection
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	))
	if err != nil {
		logger.Log.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	// NATS connection
	nc, err := natsutils.ConnectToNATS(cfg.NATSURL)
	if err != nil {
		logger.Log.Fatal("Failed to connect to NATS", zap.Error(err))
	}
	defer nc.Close()

	// Create gRPC server
	grpcServer := grpc.NewServer()
	mangaService := service.NewMangaService(db, nc)
	proto.RegisterMangaServiceServer(grpcServer, mangaService)

	// Start server
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		logger.Log.Fatal("Failed to listen", zap.Error(err))
	}

	go func() {
		logger.Log.Info("Starting gRPC server", zap.String("port", cfg.GRPCPort))
		if err := grpcServer.Serve(lis); err != nil {
			logger.Log.Fatal("Failed to serve gRPC", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("Shutting down server...")
	grpcServer.GracefulStop()
	logger.Log.Info("Server stopped")
}
