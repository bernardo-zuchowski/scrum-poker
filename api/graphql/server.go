package graphqlserver

import (
	"log"
	"multipoker/api/graphql/generated"
	gresolvers "multipoker/api/graphql/resolvers"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func Start(port string) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &gresolvers.Resolver{}}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost%s/playground for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
