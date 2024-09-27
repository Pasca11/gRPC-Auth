package service

import (
	"context"
	"fmt"
	"github.com/Pasca11/gRPC-Auth/internal/repository/postgres"
	"github.com/Pasca11/gRPC-Auth/models"
	authv1 "github.com/Pasca11/gRPC-Auth/proto/gen"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

type Service struct {
	DB *postgres.Database
	authv1.UnimplementedAuthServer
}

func RegisterServer(s *grpc.Server, db *postgres.Database) {
	authv1.RegisterAuthServer(s, &Service{
		DB: db,
	})
}

func (s *Service) Register(ctx context.Context, req *authv1.RegisterRequest) (*authv1.RegisterResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.Errorf("userService.Register.Hash: %w", err)
	}
	req.Password = string(hash)
	err = s.DB.CreateUser(&models.User{
		Username: req.Username,
		Password: req.Password,
		Role:     "user",
	})
	return &authv1.RegisterResponse{
		UserId: 0, //TODO add RETURNING in query
	}, err
}

func (s *Service) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	saved, err := s.DB.GetUserByUsername(req.Username)
	if err != nil {
		return nil, fmt.Errorf("userService.Login: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(saved.Password), []byte(req.Password))
	if err != nil {
		return nil, fmt.Errorf("userService.Login: invalid password")
	}

	token, err := createToken(saved)
	if err != nil {
		return nil, fmt.Errorf("userService.Login: %w", err)
	}

	return &authv1.LoginResponse{Token: token}, nil
}

func (s *Service) IsAdmin(ctx context.Context, req *authv1.IsAdminRequest) (*authv1.IsAdminResponse, error) {
	return nil, nil
}
