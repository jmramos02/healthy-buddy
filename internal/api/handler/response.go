package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jmramos02/healthybuddy/internal/utility"
)

type HttpErrorResponse struct {
	Errors  map[string]string `json:"error,omitempty"`
	Message string
}

func render(data interface{}, err error, context *gin.Context) {
	if err != nil {
		if serr, ok := err.(*utility.HttpError); ok {
			errorResponse := HttpErrorResponse{
				Errors:  serr.Errors,
				Message: serr.Message,
			}
			context.AbortWithStatusJSON(serr.StatusCode, errorResponse)
			return
		}
	}

	context.JSON(200, data)
	return
}
