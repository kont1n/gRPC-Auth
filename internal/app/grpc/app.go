package grpcapp

import (
	authgrpc "gRPC-Auth/internal/grpc/auth"
	"google.golang.org/grpc"
	"log/slog"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *App {

	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer)

	return &App{
		log:        log,
		port:       port,
		gRPCServer: gRPCServer,
	}
}
