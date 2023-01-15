package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Fatalf("%s\n", message)
	c.AbortWithStatusJSON(statusCode, error{Message: message})
}
