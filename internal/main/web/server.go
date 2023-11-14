package web

import (
	"database/sql"
	"net/http"

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

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	routes(s.Router, s.DB)
	http.ListenAndServe(s.WebServerPort, s.Router)
}
