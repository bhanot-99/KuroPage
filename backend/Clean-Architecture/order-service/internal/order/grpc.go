package order

import (
	"context"
	"order-service/models"

	pb "github.com/elliaaan/proto-gen/pb/order/github.com/elliaaan/proto-gen/pb/order"
)

type GRPCServer struct {
	pb.UnimplementedOrderServiceServer
	Service *Service
}

func (s *GRPCServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.OrderResponse, error) {
	order := &models.Order{
		UserID:     uint(req.Order.UserId),
		ProductID:  uint(req.Order.ProductId),
		Quantity:   int(req.Order.Quantity),
		TotalPrice: req.Order.TotalPrice,
		Status:     req.Order.Status,
	}

	if err := s.Service.CreateOrder(order); err != nil {
		return nil, err
	}
	return &pb.OrderResponse{Order: mapToProto(order)}, nil
}

func (s *GRPCServer) GetOrderByID(ctx context.Context, req *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	order, err := s.Service.GetOrderByID(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{Order: mapToProto(order)}, nil
}

func (s *GRPCServer) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.Service.ListOrders()
	if err != nil {
		return nil, err
	}

	var protoOrders []*pb.Order
	for _, o := range orders {
		protoOrders = append(protoOrders, mapToProto(&o))
	}
	return &pb.ListOrdersResponse{Orders: protoOrders}, nil
}

func (s *GRPCServer) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.OrderResponse, error) {
	order := &models.Order{
		ID:         uint(req.Order.Id),
		UserID:     uint(req.Order.UserId),
		ProductID:  uint(req.Order.ProductId),
		Quantity:   int(req.Order.Quantity),
		TotalPrice: req.Order.TotalPrice,
		Status:     req.Order.Status,
	}

	if err := s.Service.UpdateOrder(order); err != nil {
		return nil, err
	}
	return &pb.OrderResponse{Order: mapToProto(order)}, nil
}

func (s *GRPCServer) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.Empty, error) {
	if err := s.Service.DeleteOrder(uint(req.Id)); err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func mapToProto(o *models.Order) *pb.Order {
	return &pb.Order{
		Id:         uint64(o.ID),
		UserId:     uint64(o.UserID),
		ProductId:  uint64(o.ProductID),
		Quantity:   uint32(o.Quantity),
		TotalPrice: o.TotalPrice,
		Status:     o.Status,
	}
}
