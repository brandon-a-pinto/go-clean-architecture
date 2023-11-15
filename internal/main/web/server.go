package web

import (
	"net/http"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/handler"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/factory"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(port string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		WebServerPort: port,
	}
}

func routes(router chi.Router) {
	createUserHandler := handler.NewCreateUserHandler(*factory.CreateUserFactory())

	router.Post("/users", createUserHandler.CreateUser)
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	routes(s.Router)
	http.ListenAndServe(s.WebServerPort, s.Router)
}
