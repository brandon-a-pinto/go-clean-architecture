package web

import (
	"net/http"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/main/factory"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/presentation/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	swagger "github.com/swaggo/http-swagger"
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
	userHandler := handler.NewUserHandler(
		*factory.CreateUserFactory(),
		*factory.AuthenticateUserFactory(),
	)

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/auth", userHandler.AuthenticateUser)

	router.Get("/docs/*", swagger.Handler(swagger.URL("http://localhost"+s.WebServerPort+"/docs/doc.json")))
}

func (s *WebServer) Start() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	s.routes(router)

	http.ListenAndServe(s.WebServerPort, router)
}
