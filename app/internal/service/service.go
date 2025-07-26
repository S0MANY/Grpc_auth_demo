package service

import (
	"context"
	"grpcAuth/internal/repository"
	pb "grpcAuth/pkg/grpc"
	"grpcAuth/pkg/jwt"
	"grpcAuth/pkg/logger"
)

type AuthServiceInterface interface {
	Register(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error)
	Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	RefreshToken(_ context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error)
}

type Service struct {
	Auth AuthServiceInterface
}

func NewService(repo *repository.Repository, jwt *jwt.JWTManager, log *logger.Logger) *Service {
	return &Service{
		Auth: NewAuthServer(repo, jwt, log),
	}
}

func (s *Service) GetGRPCServer() pb.AuthServiceServer {
	return s.Auth.(pb.AuthServiceServer)
}
