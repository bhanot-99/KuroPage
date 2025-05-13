package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bhanot-99/KuroPage/Backend/pkg/config"
	"github.com/bhanot-99/KuroPage/Backend/pkg/logger"
	"github.com/bhanot-99/KuroPage/Backend/pkg/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger.Init()
	defer logger.Sync()

	cfg := config.Load()

	// Create gRPC clients
	userConn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Log.Fatal("Failed to connect to user service", zap.Error(err))
	}
	defer userConn.Close()
	userClient := proto.NewUserServiceClient(userConn)

	mangaConn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.MangaServiceHost, cfg.MangaServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Log.Fatal("Failed to connect to manga service", zap.Error(err))
	}
	defer mangaConn.Close()
	mangaClient := proto.NewMangaServiceClient(mangaConn)

	orderConn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.OrderServiceHost, cfg.OrderServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Log.Fatal("Failed to connect to order service", zap.Error(err))
	}
	defer orderConn.Close()
	orderClient := proto.NewOrderServiceClient(orderConn)

	// Create HTTP server with gRPC gateway
	mux := runtime.NewServeMux()
	ctx := context.Background()

	// Register handlers
	if err := proto.RegisterUserServiceHandlerClient(ctx, mux, userClient); err != nil {
		logger.Log.Fatal("Failed to register user service handler", zap.Error(err))
	}

	if err := proto.RegisterMangaServiceHandlerClient(ctx, mux, mangaClient); err != nil {
		logger.Log.Fatal("Failed to register manga service handler", zap.Error(err))
	}

	if err := proto.RegisterOrderServiceHandlerClient(ctx, mux, orderClient); err != nil {
		logger.Log.Fatal("Failed to register order service handler", zap.Error(err))
	}

	// Add middleware for logging, auth, etc.
	handler := addMiddleware(mux)

	// Start server
	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: handler,
	}

	go func() {
		logger.Log.Info("Starting HTTP server", zap.String("port", cfg.HTTPPort))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Fatal("Failed to serve HTTP", zap.Error(err))
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Log.Info("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Error("Failed to shutdown server gracefully", zap.Error(err))
	}
	logger.Log.Info("Server stopped")
}

func addMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log request
		logger.Log.Info("Incoming request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("remote_addr", r.RemoteAddr))

		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// TODO: Add authentication middleware
		// For now, just pass through
		next.ServeHTTP(w, r)
	})
}
