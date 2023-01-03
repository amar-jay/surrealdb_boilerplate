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


	// database connection
	dbCtx := context.Background()
	Db, err := dbConn(&dbCtx)
	if err != nil {
		log.Fatal(err)
	}

	defer Db.Close()
	// test database connection
	//Db.Info(&dbCtx)
	Db.Use(&dbCtx, "test", "test")
	Db.Select(&dbCtx, "1")

  /*
	app.GET("/api", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})*/

  // TODO: fix gin integration with gqlgen
	//app.GET("/graphql", playgroundHandler())
	//app.GET("/query", graphqlHandler())
	http.Handle("/graphql", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", 
	  handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		  Resolvers:  &graph.Resolver{Db: Db},
		  Directives: graph.DirectiveRoot{},
		  Complexity: graph.ComplexityRoot{},
	    })))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	//log.Fatal(app.Run(":"+port))
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
