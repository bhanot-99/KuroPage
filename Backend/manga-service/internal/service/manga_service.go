package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/bhanot-99/KuroPage/Backend/manga-service/internal/model"
	"github.com/bhanot-99/KuroPage/Backend/manga-service/internal/repository"
	"github.com/bhanot-99/KuroPage/Backend/pkg/logger"
	"github.com/bhanot-99/KuroPage/Backend/pkg/natsutils"
	"github.com/bhanot-99/KuroPage/Backend/pkg/proto"
	"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

type MangaService struct {
	repo repository.MangaRepository
	nc   *nats.Conn
	proto.UnimplementedMangaServiceServer
}

func NewMangaService(db *sqlx.DB, nc *nats.Conn) *MangaService {
	svc := &MangaService{
		repo: repository.NewMangaRepository(db),
		nc:   nc,
	}

	// Subscribe to order.created events
	_, err := natsutils.SubscribeToTopic(nc, "order.created", svc.handleOrderCreated)
	if err != nil {
		logger.Log.Fatal("Failed to subscribe to order.created topic", zap.Error(err))
	}

	return svc
}

func (s *MangaService) AddManga(ctx context.Context, req *proto.AddMangaRequest) (*proto.AddMangaResponse, error) {
	manga := &model.Manga{
		Title:       req.Manga.Title,
		Author:      req.Manga.Author,
		Description: req.Manga.Description,
		Price:       req.Manga.Price,
		Stock:       int(req.Manga.Stock),
	}

	if err := s.repo.CreateManga(ctx, manga); err != nil {
		logger.Log.Error("Failed to create manga", zap.Error(err))
		return nil, errors.New("internal server error")
	}

	return &proto.AddMangaResponse{MangaId: manga.ID}, nil
}

func (s *MangaService) EditManga(ctx context.Context, req *proto.EditMangaRequest) (*proto.EditMangaResponse, error) {
	manga := &model.Manga{
		ID:          req.Manga.Id,
		Title:       req.Manga.Title,
		Author:      req.Manga.Author,
		Description: req.Manga.Description,
		Price:       req.Manga.Price,
		Stock:       int(req.Manga.Stock),
	}

	if err := s.repo.UpdateManga(ctx, manga); err != nil {
		logger.Log.Error("Failed to update manga", zap.Error(err))
		return nil, errors.New("internal server error")
	}

	return &proto.EditMangaResponse{Success: true}, nil
}

func (s *MangaService) ListMangas(ctx context.Context, req *proto.ListMangasRequest) (*proto.ListMangasResponse, error) {
	mangas, total, err := s.repo.ListMangas(ctx, int(req.Page), int(req.Limit))
	if err != nil {
		logger.Log.Error("Failed to list mangas", zap.Error(err))
		return nil, errors.New("internal server error")
	}

	var protoMangas []*proto.Manga
	for _, m := range mangas {
		protoMangas = append(protoMangas, &proto.Manga{
			Id:          m.ID,
			Title:       m.Title,
			Author:      m.Author,
			Description: m.Description,
			Price:       m.Price,
			Stock:       int32(m.Stock),
		})
	}

	return &proto.ListMangasResponse{
		Mangas: protoMangas,
		Total:  int32(total),
	}, nil
}

func (s *MangaService) CheckMangaAvailability(ctx context.Context, req *proto.CheckMangaAvailabilityRequest) (*proto.CheckMangaAvailabilityResponse, error) {
	manga, err := s.repo.GetMangaByID(ctx, req.MangaId)
	if err != nil {
		logger.Log.Error("Failed to get manga by ID", zap.Error(err))
		return nil, errors.New("internal server error")
	}
	if manga == nil {
		return &proto.CheckMangaAvailabilityResponse{
			Available: false,
			Stock:     0,
		}, nil
	}

	available := manga.Stock >= int(req.Quantity)
	return &proto.CheckMangaAvailabilityResponse{
		Available: available,
		Stock:     int32(manga.Stock),
	}, nil
}

func (s *MangaService) handleOrderCreated(msg *nats.Msg) {
	var order struct {
		Items []struct {
			MangaID  string `json:"manga_id"`
			Quantity int    `json:"quantity"`
		} `json:"items"`
	}

	if err := json.Unmarshal(msg.Data, &order); err != nil {
		logger.Log.Error("Failed to unmarshal order.created message", zap.Error(err))
		return
	}

	ctx := context.Background()
	for _, item := range order.Items {
		if err := s.repo.UpdateStock(ctx, item.MangaID, item.Quantity); err != nil {
			logger.Log.Error("Failed to update manga stock",
				zap.String("manga_id", item.MangaID),
				zap.Error(err))
			// In a real application, you might want to implement a retry mechanism
			// or send a notification about the failure
		}
	}
}
