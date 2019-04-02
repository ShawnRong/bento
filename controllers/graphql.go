package controllers

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/ShawnRong/bento/graphql/generated"
	"github.com/ShawnRong/bento/graphql/resolver"
	"github.com/gin-gonic/gin"
)

// Defining the Graphql handler
func GraphqlHandler() gin.HandlerFunc {
	c := generated.Config{
		Resolvers: &resolver.Resolver{},
	}
	//@TODO implement Directives

	h := handler.GraphQL(generated.NewExecutableSchema(c))

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
