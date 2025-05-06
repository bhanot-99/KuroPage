package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/bhanot-99/KuroPage/Backend/order-service/internal/model"
	"github.com/bhanot-99/KuroPage/Backend/order-service/internal/repository"
	"github.com/bhanot-99/KuroPage/Backend/pkg/logger"
	"github.com/bhanot-99/KuroPage/Backend/pkg/natsutils"
	"github.com/bhanot-99/KuroPage/Backend/pkg/proto"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type OrderService struct {
	repo        repository.OrderRepository
	nc          *nats.Conn
	mangaClient proto.MangaServiceClient
	proto.UnimplementedOrderServiceServer
}

func NewOrderService(db *sqlx.DB, nc *nats.Conn, mangaClient proto.MangaServiceClient) *OrderService {
	return &OrderService{
		repo:        repository.NewOrderRepository(db),
		nc:          nc,
		mangaClient: mangaClient,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req *proto.CreateOrderRequest) (*proto.CreateOrderResponse, error) {
	// Validate manga availability
	for _, item := range req.Items {
		resp, err := s.mangaClient.CheckMangaAvailability(ctx, &proto.CheckMangaAvailabilityRequest{
			MangaId:  item.MangaId,
			Quantity: item.Quantity,
		})
		if err != nil {
			logger.Log.Error("Failed to check manga availability",
				zap.String("manga_id", item.MangaId),
				zap.Error(err))
			return nil, errors.New("failed to validate order")
		}
		if !resp.Available {
			return nil, errors.New("manga not available or insufficient stock")
		}
	}

	// Calculate total amount
	var totalAmount float64
	var orderItems []*model.OrderItem
	for _, item := range req.Items {
		totalAmount += item.UnitPrice * float64(item.Quantity)
		orderItems = append(orderItems, &model.OrderItem{
			MangaID:   item.MangaId,
			Quantity:  int(item.Quantity),
			UnitPrice: item.UnitPrice,
		})
	}

	// Create order
	order := &model.Order{
		UserID:      req.UserId,
		TotalAmount: totalAmount,
		Status:      "pending",
	}

	if err := s.repo.CreateOrder(ctx, order, orderItems); err != nil {
		logger.Log.Error("Failed to create order", zap.Error(err))
		return nil, errors.New("failed to create order")
	}

	// Publish order.created event
	orderCreatedEvent := struct {
		OrderID string `json:"order_id"`
		UserID  string `json:"user_id"`
		Items   []struct {
			MangaID  string `json:"manga_id"`
			Quantity int    `json:"quantity"`
		} `json:"items"`
	}{
		OrderID: order.ID,
		UserID:  req.UserId,
	}

	for _, item := range req.Items {
		orderCreatedEvent.Items = append(orderCreatedEvent.Items, struct {
			MangaID  string `json:"manga_id"`
			Quantity int    `json:"quantity"`
		}{
			MangaID:  item.MangaId,
			Quantity: int(item.Quantity),
		})
	}

	eventData, err := json.Marshal(orderCreatedEvent)
	if err != nil {
		logger.Log.Error("Failed to marshal order.created event", zap.Error(err))
		// We still return success because the order was created, just the event failed
		return &proto.CreateOrderResponse{
			OrderId:     order.ID,
			TotalAmount: totalAmount,
		}, nil
	}

	if err := natsutils.PublishMessage(s.nc, "order.created", eventData); err != nil {
		logger.Log.Error("Failed to publish order.created event", zap.Error(err))
		// Same as above - order was created, just event failed
	}

	return &proto.CreateOrderResponse{
		OrderId:     order.ID,
		TotalAmount: totalAmount,
	}, nil
}
