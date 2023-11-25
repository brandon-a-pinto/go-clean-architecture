package graph

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

type GraphQLServer struct {
	Port string
}

func NewGraphQLServer(port string) *GraphQLServer {
	return &GraphQLServer{
		Port: port,
	}
}

func (s *GraphQLServer) Start() {
	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	http.ListenAndServe(":"+s.Port, nil)
}
