package web

import (
	"database/sql"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/handler"
	"github.com/go-chi/chi/v5"
)

func routes(router chi.Router, db *sql.DB) {
	userHandler := handler.NewUserHandler(db)

	router.Post("/users", userHandler.Create)
}
