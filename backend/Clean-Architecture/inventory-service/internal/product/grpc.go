package product

import (
	"context"
	"inventory-service/models"

	pb "github.com/elliaaan/proto-gen/pb/inventory/github.com/elliaaan/proto-gen/pb/inventory"
)

type GRPCServer struct {
	pb.UnimplementedInventoryServiceServer
	Service *Service
}

func (s *GRPCServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	product := &models.Product{
		Name:     req.Product.Name,
		Category: req.Product.Category,
		Price:    float64(req.Product.Price),
		Stock:    int(req.Product.Stock),
	}

	if err := s.Service.CreateProduct(product); err != nil {
		return nil, err
	}

	return &pb.ProductResponse{Product: mapToProto(product)}, nil
}

func (s *GRPCServer) GetProductByID(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	product, err := s.Service.GetProductByID(uint(req.Id))
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{Product: mapToProto(product)}, nil
}

func (s *GRPCServer) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	products, err := s.Service.GetProducts()
	if err != nil {
		return nil, err
	}

	var protoProducts []*pb.Product
	for _, p := range products {
		protoProducts = append(protoProducts, mapToProto(&p))
	}

	return &pb.ListProductsResponse{Products: protoProducts}, nil
}

func (s *GRPCServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.ProductResponse, error) {
	product := &models.Product{
		ID:       uint(req.Product.Id),
		Name:     req.Product.Name,
		Category: req.Product.Category,
		Price:    float64(req.Product.Price),
		Stock:    int(req.Product.Stock),
	}

	updates := map[string]interface{}{
		"name":     product.Name,
		"category": product.Category,
		"price":    product.Price,
		"stock":    product.Stock,
	}

	err := s.Service.UpdateProduct(product.ID, updates)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{Product: mapToProto(product)}, nil
}

func (s *GRPCServer) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.Empty, error) {
	err := s.Service.DeleteProduct(uint(req.Id))
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func mapToProto(p *models.Product) *pb.Product {
	return &pb.Product{
		Id:       uint64(p.ID),
		Name:     p.Name,
		Category: p.Category,
		Price:    float64(p.Price),
		Stock:    uint32(p.Stock),
	}
}
