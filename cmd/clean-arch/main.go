package main

import (
	"fmt"

	"github.com/brandon-a-pinto/go-clean-architecture/configs"
	_ "github.com/brandon-a-pinto/go-clean-architecture/docs"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/grpc"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/web"
	"github.com/brandon-a-pinto/go-clean-architecture/pkg/infra/database"

	_ "github.com/lib/pq"
)

// @title           Go Clean Architecture
// @version         0.0.1
// @description     Go (Golang) Clean Architecture API project with gRPC, PostgreSQL, Docker and more.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Brandon Pinto
// @contact.email  brandon.amaral9658@gmail.com

// @license.name  GPL-3.0
// @license.url   https://www.gnu.org/licenses/gpl-3.0.en.html

// @host      localhost:3000
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func main() {
	config := configs.LoadConfig()

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
