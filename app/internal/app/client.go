package app

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "grpcAuth/pkg/grpc"
	"grpcAuth/pkg/logger"
	"time"
)

func ImitateClient() {
	log := logger.NewLogger()
	ctx := context.Background()

	conn, err := grpc.NewClient("localhost:228", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error start Client: ", err.Error())
	}
	defer conn.Close()

	authClient := pb.NewAuthServiceClient(conn)

	response, err := authClient.Register(ctx, &pb.RegisterRequest{
		Id:               "2",
		Username:         "Andrew",
		Password:         "andrew332",
		RepeatedPassword: "andrew332",
		CreatedAt:        timestamppb.New(time.Now()),
	})
	if err != nil {
		log.Fatal("Err: %s", err.Error())
	}

	log.Info("Response: %s", response.String())

	response1, err := authClient.Login(ctx, &pb.LoginRequest{
		Username: "Andrew",
		Password: "Andrew",
	})

	if err != nil {
		log.Fatal("Err: %s", err.Error())
	}

	log.Info("Access Token: %s | Refresh Token: %s", response1.GetAccessToken(), response1.GetRefreshToken())
}
