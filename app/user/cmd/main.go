package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"thesis/be/app/user/api"

	"google.golang.org/grpc"
)

type UserServer struct {
	api.UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(ctx context.Context, in *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	fmt.Print("helllllllllll")
	return &api.CreateUserResponse{
		Code:    "3",
		Message: "à haha",
	}, nil
}
func (s *UserServer) GetUser(ctx context.Context, in *api.GetUserRequest) (*api.GetUserResponse, error) {
	fmt.Print("helllllllllll")
	return &api.GetUserResponse{
		Code:    "3",
		Message: "à haha",
	}, nil
}

func main() {

	fmt.Println("Go gRPC Beginners Tutorial!")

	lis, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Print("??")

	s := UserServer{}

	grpcServer := grpc.NewServer()

	api.RegisterUserServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
	fmt.Print("sssss??")

}
