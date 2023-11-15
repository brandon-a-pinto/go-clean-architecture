package grpc

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/service"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/factory"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	DB             *sql.DB
	GRPCServerPort string
}

func NewGRPCServer(port string, db *sql.DB) *GRPCServer {
	return &GRPCServer{
		DB:             db,
		GRPCServerPort: port,
	}
}

func services(server grpc.ServiceRegistrar, db *sql.DB) {
	createUserService := service.NewCreateUserService(db, *factory.CreateUserFactory(db))

	pb.RegisterUserServiceServer(server, createUserService)
}

func (s *GRPCServer) Start() {
	server := grpc.NewServer()

	services(server, s.DB)
	reflection.Register(server)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", s.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	server.Serve(listen)
}
