package inventory

import (
	"context"
	"log"
	"time"

	pb "github.com/elliaaan/proto-gen/pb/inventory/github.com/elliaaan/proto-gen/pb/inventory"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	Client pb.InventoryServiceClient
}

func NewInventoryClient(address string) *Client {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to inventory service: %v", err)
	}
	client := pb.NewInventoryServiceClient(conn)
	return &Client{conn: conn, Client: client}
}

func (c *Client) CreateProduct(product *pb.Product) (*pb.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return c.Client.CreateProduct(ctx, &pb.CreateProductRequest{Product: product})
}

func (c *Client) GetProductByID(id uint64) (*pb.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return c.Client.GetProductByID(ctx, &pb.GetProductRequest{Id: id})
}

func (c *Client) UpdateProduct(product *pb.Product) (*pb.ProductResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return c.Client.UpdateProduct(ctx, &pb.UpdateProductRequest{Product: product})
}

func (c *Client) DeleteProduct(id uint64) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return c.Client.DeleteProduct(ctx, &pb.DeleteProductRequest{Id: id})
}

func (c *Client) ListProducts() ([]*pb.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := c.Client.ListProducts(ctx, &pb.ListProductsRequest{})
	if err != nil {
		return nil, err
	}

	return res.Products, nil
}
