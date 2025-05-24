package statistics

import (
	"context"

	statisticspb "github.com/elliaaan/proto-gen/pb/statistics"
)

type GRPCServer struct {
	statisticspb.UnimplementedStatisticsServiceServer
	Service *Service
}

func (s *GRPCServer) GetUserStatistics(ctx context.Context, req *statisticspb.UserStatisticsRequest) (*statisticspb.UserStatisticsResponse, error) {
	stats, err := s.Service.GetUserStatistics(req.UserId)
	if err != nil {
		return nil, err
	}

	return &statisticspb.UserStatisticsResponse{
		UserId:        req.UserId,
		TotalEvents:   stats.Total,
		OrderEvents:   stats.Orders,
		ProductEvents: stats.Products,
	}, nil
}
func (s *GRPCServer) GetUserOrdersStatistics(ctx context.Context, req *statisticspb.UserOrderStatisticsRequest) (*statisticspb.UserOrderStatisticsResponse, error) {
	count, err := s.Service.GetOrderCountByUser(req.UserId)
	if err != nil {
		return nil, err
	}

	return &statisticspb.UserOrderStatisticsResponse{
		UserId:      req.UserId,
		TotalOrders: count,
	}, nil
}
