package web

import (
	"net/http"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/handler"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/factory"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	WebServerPort string
}

func NewWebServer(port string) *WebServer {
	return &WebServer{
		WebServerPort: port,
	}
}

func (s *WebServer) routes(router chi.Router) {
	createUserHandler := handler.NewCreateUserHandler(*factory.CreateUserFactory())
	authenticateUserHandler := handler.NewAuthenticateUserHandler(*factory.AuthenticateUserFactory())

	router.Post("/users", createUserHandler.CreateUser)
	router.Post("/users/auth", authenticateUserHandler.AuthenticateUser)
}

func (s *WebServer) Start() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	s.routes(router)

	http.ListenAndServe(s.WebServerPort, router)
}
