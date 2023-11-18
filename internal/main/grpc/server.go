package grpc

import (
	"fmt"
	"net"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/service"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/factory"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	GRPCServerPort string
}

func NewGRPCServer(port string) *GRPCServer {
	return &GRPCServer{
		GRPCServerPort: port,
	}
}

func services(server grpc.ServiceRegistrar) {
	userService := service.NewUserService(
		*factory.CreateUserFactory(),
		*factory.AuthenticateUserFactory(),
	)

	pb.RegisterUserServiceServer(server, userService)
}

func (s *GRPCServer) Start() {
	server := grpc.NewServer()

	services(server)
	reflection.Register(server)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	server.Serve(listen)
}
