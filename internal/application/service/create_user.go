package service

import (
	"context"
	"time"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/usecase"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/dto"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/grpc/pb"
)

type CreateUserService struct {
	pb.UnimplementedUserServiceServer
	CreateUserUsecase usecase.CreateUserUsecase
}

func NewCreateUserService(createUserUsecase usecase.CreateUserUsecase) *CreateUserService {
	return &CreateUserService{
		CreateUserUsecase: createUserUsecase,
	}
}

func (s *CreateUserService) CreateUser(ctx context.Context, input *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
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
