package main

import (
	"database/sql"
	"fmt"

	"github.com/brandon-a-pinto/go-clean-architecture/configs"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/web"

	_ "github.com/lib/pq"
)

func main() {
	// Configuration
	config := configs.LoadConfigDocker()

	// Database
	db, err := sql.Open(config.DBDriver, fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", config.DBDriver, config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Web Server
	server := web.NewWebServer(":"+config.WebServerPort, db)
	fmt.Println("Starting web server on port", config.WebServerPort)
	server.Start()
}
