package app

import (
	"google.golang.org/grpc"
	"grpcAuth/internal/repository"
	"grpcAuth/internal/service"
	pb "grpcAuth/pkg/grpc"
	"grpcAuth/pkg/jwt"
	"grpcAuth/pkg/logger"

	"net"
)

func StartServer() {
	log := logger.NewLogger()
	repo := repository.NewReposiitory()
	jwtManager := jwt.NewJWTManager("123", "456", 300, 60000)
	serv := service.NewService(repo, jwtManager, log)

	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, serv.GetGRPCServer())

	listener, err := net.Listen("tcp", ":228")
	if err != nil {
		log.Fatal("Error start listener: %s", err.Error())
	}

	log.Info("GRPC Ready amd listen on 228 port")
	if err := s.Serve(listener); err != nil {
		log.Fatal("Error Start Server With GRPC: %s", err.Error())
	}
}
