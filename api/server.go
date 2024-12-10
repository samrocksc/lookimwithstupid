package main

import (
	db "fun/database"
	"fun/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	mongoUri := "mongodb://localhost:27017"
	dbName := "a-db"

	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	db, err := db.ConnectToMongoDB(mongoUri, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	if db == nil {
		log.Fatalf("db is nil")
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
