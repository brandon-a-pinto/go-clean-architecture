package web

import (
	"database/sql"
	"net/http"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/handler"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/factory"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	DB            *sql.DB
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(port string, db *sql.DB) *WebServer {
	return &WebServer{
		DB:            db,
		Router:        chi.NewRouter(),
		WebServerPort: port,
	}
}

func routes(router chi.Router, db *sql.DB) {
	createUserHandler := handler.NewCreateUserHandler(db, *factory.CreateUserFactory(db))

	router.Post("/users", createUserHandler.CreateUser)
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	routes(s.Router, s.DB)
	http.ListenAndServe(s.WebServerPort, s.Router)
}
