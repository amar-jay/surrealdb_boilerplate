package graph

import (
	"github.com/99designs/gqlgen/graphql/playground"
_	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
)

// PlaygroundHandler Defines the Playground handler to expose our playground
func PlaygroundHandler(path string) gin.HandlerFunc {
	h := playground.Handler("GraphQL Playground", path)
	return func(c *gin.Context) { h.ServeHTTP(c.Writer, c.Request) }
}
