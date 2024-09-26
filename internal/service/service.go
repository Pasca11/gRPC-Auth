package service

import (
	"context"
	authv1 "github.com/Pasca11/authservice/proto/gen"
	"google.golang.org/grpc"
)

type Service struct {
	authv1.UnimplementedAuthServer
}

func RegisterServer(s *grpc.Server) {
	authv1.RegisterAuthServer(s, &Service{})
}

func Register(ctx context.Context, request authv1.RegisterRequest) (*authv1.RegisterResponse, error) {

}
