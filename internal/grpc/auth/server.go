package auth

import (
	"context"

	"google.golang.org/grpc"

	ssov1 "github.com/kont1n/Protobuf/gen/go/sso"
)

type serverAPI struct {
	ssov1.UnimplementedAuthServiceServer
}

func Register(gRPC *grpc.Server) {
	ssov1.RegisterAuthServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Register(
	ctx context.Context,
	request *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	panic("implement me")
}

func (s *serverAPI) Login(
	ctx context.Context,
	request *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	panic("implement me")
}

func (s *serverAPI) IsAdmin(
	ctx context.Context,
	request *ssov1.IsAdminRequest,
) (*ssov1.IsAdminResponse, error) {
	panic("implement me")
}
