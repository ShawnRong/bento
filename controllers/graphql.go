package controllers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/ShawnRong/bento/graphql/generated"
	"github.com/ShawnRong/bento/graphql/resolver"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func GraphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{
		Resolvers: &resolver.Resolver{},
	}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func PlaygroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
