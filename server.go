// server.go
package main

import (
	"log"
	"net/http"

	"graphQLGEN/graph"
	"graphQLGEN/utils"

	"github.com/99designs/gqlgen/handler"
)

func main() {
	// Set up GORM database connection
	utils.SetupDB()

	// Get the GORM database instance
	db := utils.GetDB()

	// Initialize resolver with the database connection
	resolver := graph.NewResolver(db)

	// Create GraphQL server
	http.Handle("/query", handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: resolver})))

	// Start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
