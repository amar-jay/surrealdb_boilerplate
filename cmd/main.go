package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/amar-jay/surreal/graph"
	"github.com/gin-gonic/gin"
)

var (
	app         = gin.Default()
	defaultPort = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	dbCtx := context.Background()
	// database connection
	db, err := dbConn(&dbCtx)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// test database connection
	db.Info(&dbCtx)
	db.Select(&dbCtx, "1")

	app.GET("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
