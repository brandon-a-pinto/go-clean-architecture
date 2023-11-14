package main

import (
	"database/sql"
	"fmt"

	"github.com/brandon-a-pinto/go-clean-architecture/configs"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/web"
	"github.com/sagikazarmark/slog-shim"

	_ "github.com/lib/pq"
)

func main() {
	config := configs.LoadConfigDocker()

	db, err := sql.Open(config.DBDriver, fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", config.DBDriver, config.DBUsername, config.DBPassword, config.DBHost, config.DBPort, config.DBName))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	server := web.NewWebServer(":" + config.WebServerPort)
	slog.Info("Starting server...")
	server.Start()
}
