package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/handler"
)

func GraphQLHandler(config *handler.Config) gin.HandlerFunc {
	h := handler.New(config)

	return func(context *gin.Context) {
		h.ServeHTTP(context.Writer, context.Request)
	}
}
