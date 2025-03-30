package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Jeanga7/graphql-go-api.git/db"
	"github.com/Jeanga7/graphql-go-api.git/graph"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	// Connect to the database and migrate the schema
	dbInstance := db.ConnectDatabase()
	db.MigrateDatabase()
	log.Println("Database connected and migrations completed")

	// Retrieve the port from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// Initialize the resolver
	resolver := graph.NewResolver(dbInstance)

	
	// Create the GraphQL server
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	// Add transports and middleware
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	// Set up the HTTP routes for the GraphQL Playground and the query handler
	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	// Start the server
	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
