package service

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"grpcAuth/internal/repository"
	pb "grpcAuth/pkg/grpc"
	"grpcAuth/pkg/jwt"
	"grpcAuth/pkg/logger"
	"time"
)

type AuthServer struct {
	Log  *logger.Logger
	Repo *repository.Repository
	JWT  *jwt.JWTManager
	pb.UnimplementedAuthServiceServer
}

func NewAuthServer(repo *repository.Repository, jwt *jwt.JWTManager, log *logger.Logger) *AuthServer {
	return &AuthServer{
		Repo: repo,
		JWT:  jwt,
		Log:  log,
	}
}

func (au *AuthServer) Register(_ context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	Id := uuid.New().String()
	username := req.GetUsername()
	password := req.GetPassword()
	repeatedPassword := req.GetRepeatedPassword()
	createdAt := req.GetCreatedAt()

	if username == "" {
		return nil, status.Error(codes.InvalidArgument, "Name can't be Empty")
	}

	if createdAt.AsTime().After(time.Now()) {
		return nil, status.Error(codes.InvalidArgument, "Invalid created at field")
	}

	if len(password) < 5 {
		return nil, status.Error(codes.InvalidArgument, "Invalid password")
	}

	if password != repeatedPassword {
		return nil, status.Error(codes.InvalidArgument, "Different passwords")
	}

	isExists := au.Repo.Users.FindWithUsername(username)

	if isExists {
		return nil, status.Error(codes.InvalidArgument, "Username already exists.")
	} else {
		return &pb.RegisterResponse{Id: Id}, nil
	}
}

func (au *AuthServer) Login(_ context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	username := req.GetUsername()
	password := req.GetPassword()

	isValid := au.Repo.Users.CheckAccessWithPassword(username, password)

	if isValid {
		accessToekn, err := au.JWT.GenerateAccessToken(username)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}

		refreshToken, err := au.JWT.GenerateRefreshToken(username)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		return &pb.LoginResponse{
			AccessToken:  accessToekn,
			RefreshToken: refreshToken,
		}, nil
	} else {
		return nil, status.Error(codes.PermissionDenied, "Invalid username or password.")
	}
}

func (au *AuthServer) RefreshToken(_ context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {

	token := req.GetRefreshToken()

	claims, err := au.JWT.VerifyRefreshToken(token)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	accessToken, err := au.JWT.GenerateAccessToken(claims.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	refreshToken, err := au.JWT.GenerateRefreshToken(claims.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
