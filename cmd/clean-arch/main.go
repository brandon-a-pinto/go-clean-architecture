package main

import (
	"fmt"

	"github.com/brandon-a-pinto/go-clean-architecture/configs"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/infra/database"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/grpc"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/web"

	_ "github.com/lib/pq"
)

func main() {
	// Configuration
	config := configs.LoadConfigDocker()

	// Database
	db := database.Start(config.DBDriver, config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	defer db.Close()

	// Web Server
	server := web.NewWebServer(":" + config.WebServerPort)
	fmt.Println("Starting web server on port", config.WebServerPort)
	go server.Start()

	// gRPC Server
	grpc := grpc.NewGRPCServer(config.GRPCServerPort)
	fmt.Println("Starting gRPC server on port", config.GRPCServerPort)
	grpc.Start()
}
