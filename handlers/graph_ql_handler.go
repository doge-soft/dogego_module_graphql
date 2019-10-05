package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func GraphQLHandler(schema *graphql.Schema) gin.HandlerFunc {
	h := handler.New(&handler.Config{
		Schema:   schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return func(context *gin.Context) {
		h.ServeHTTP(context.Writer, context.Request)
	}
}
