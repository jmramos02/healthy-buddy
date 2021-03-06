package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jmramos02/healthy-buddy/internal/domain/authentication"
)

func RegisterDietitian(c *gin.Context) {
	var request authentication.RegisterRequest
	c.Bind(&request)

	ctx := c.MustGet("context").(context.Context)
	response, err := authentication.Register(ctx, request, "dietitian")

	render(response, err, c)
}
