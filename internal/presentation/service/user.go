package service

import (
	"context"
	"time"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/usecase"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/dto"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	CreateUserUsecase       usecase.CreateUserUsecase
	AuthenticateUserUsecase usecase.AuthenticateUserUsecase
}

func NewUserService(
	createUserUsecase usecase.CreateUserUsecase,
	authenticateUserUsecase usecase.AuthenticateUserUsecase,
) *UserService {
	return &UserService{
		CreateUserUsecase:       createUserUsecase,
		AuthenticateUserUsecase: authenticateUserUsecase,
	}
}

func (s *UserService) CreateUser(ctx context.Context, input *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	dto := dto.CreateUserInput{
		Email:       input.Email,
		Username:    input.Username,
		DisplayName: input.DisplayName,
		Password:    input.Password,
	}

	output, err := s.CreateUserUsecase.Exec(ctx, time.Second*5, dto)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Email:       output.Email,
		Username:    output.Username,
		DisplayName: output.DisplayName,
	}, nil
}

func (s *UserService) AuthenticateUser(ctx context.Context, input *pb.AuthenticateUserRequest) (*pb.AuthenticateUserResponse, error) {
	dto := dto.AuthenticateUserInput{
		Email:    input.Email,
		Password: input.Password,
	}

	output, err := s.AuthenticateUserUsecase.Exec(ctx, time.Second*5, dto)
	if err != nil {
		return nil, err
	}

	return &pb.AuthenticateUserResponse{
		AccessToken: output.AccessToken,
	}, nil
}
