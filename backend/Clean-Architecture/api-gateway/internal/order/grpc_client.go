package order

import (
	"context"
	"log"
	"time"

	pb "github.com/elliaaan/proto-gen/pb/order/github.com/elliaaan/proto-gen/pb/order"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	Client pb.OrderServiceClient
}

func NewOrderClient(address string) *Client {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	client := pb.NewOrderServiceClient(conn)
	return &Client{conn: conn, Client: client}
}

func (c *Client) CreateOrder(order *pb.Order) (*pb.OrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return c.Client.CreateOrder(ctx, &pb.CreateOrderRequest{Order: order})
}

func (c *Client) GetOrderByID(id uint64) (*pb.OrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return c.Client.GetOrderByID(ctx, &pb.GetOrderRequest{Id: id})
}

func (c *Client) UpdateOrder(order *pb.Order) (*pb.OrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return c.Client.UpdateOrder(ctx, &pb.UpdateOrderRequest{Order: order})
}

func (c *Client) DeleteOrder(id uint64) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return c.Client.DeleteOrder(ctx, &pb.DeleteOrderRequest{Id: id})
}

func (c *Client) ListOrders() ([]*pb.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := c.Client.ListOrders(ctx, &pb.ListOrdersRequest{})
	if err != nil {
		return nil, err
	}
	return res.Orders, nil
}
