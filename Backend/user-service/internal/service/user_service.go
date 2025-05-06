package service

import (
	"context"
	"errors"

	"github.com/bhanot-99/KuroPage/Backend/pkg/logger"
	"github.com/bhanot-99/KuroPage/Backend/pkg/proto"
	"github.com/bhanot-99/KuroPage/Backend/user-service/internal/model"
	"github.com/bhanot-99/KuroPage/Backend/user-service/internal/repository"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repository.UserRepository
	proto.UnimplementedUserServiceServer
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		repo: repository.NewUserRepository(db),
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {
	// Check if user already exists
	existingUser, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Log.Error("Failed to check user existence", zap.Error(err))
		return nil, errors.New("internal server error")
	}
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.Error("Failed to hash password", zap.Error(err))
		return nil, errors.New("internal server error")
	}

	// Create user
	user := &model.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		logger.Log.Error("Failed to create user", zap.Error(err))
		return nil, errors.New("internal server error")
	}

	return &proto.RegisterUserResponse{UserId: user.ID}, nil
}

func (s *UserService) LoginUser(ctx context.Context, req *proto.LoginUserRequest) (*proto.LoginUserResponse, error) {
	user, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Log.Error("Failed to get user by email", zap.Error(err))
		return nil, errors.New("internal server error")
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// In a real application, generate a JWT token here
	token := "generated-jwt-token"

	return &proto.LoginUserResponse{Token: token}, nil
}

func (s *UserService) GetUserProfile(ctx context.Context, req *proto.GetUserProfileRequest) (*proto.GetUserProfileResponse, error) {
	user, err := s.repo.GetUserByID(ctx, req.UserId)
	if err != nil {
		logger.Log.Error("Failed to get user by ID", zap.Error(err))
		return nil, errors.New("internal server error")
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return &proto.GetUserProfileResponse{
		UserId:   user.ID,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
