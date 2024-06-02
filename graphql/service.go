package graphql

import (
	"cmk/generated/graphql"
	"net/http"

	"encore.dev"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

//go:generate go run github.com/99designs/gqlgen generate

//encore:service
type Service struct {
	srv        *handler.Server
	playground http.Handler
}

func initService() (*Service, error) {
	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{}}),
	)
	pg := playground.Handler("GraphQL Playground", "/graphql")
	return &Service{srv: srv, playground: pg}, nil
}

//encore:api public raw path=/graphql
func (s *Service) Query(w http.ResponseWriter, req *http.Request) {
	s.srv.ServeHTTP(w, req)
}

//encore:api public raw path=/graphql/playground
func (s *Service) Playground(w http.ResponseWriter, req *http.Request) {
	// Disable playground in production
	if encore.Meta().Environment.Type == encore.EnvProduction {
		http.Error(w, "Playground disabled", http.StatusNotFound)
		return
	}

	s.playground.ServeHTTP(w, req)
}
