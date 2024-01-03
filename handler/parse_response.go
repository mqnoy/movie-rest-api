package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mqnoy/movie-rest-api/pkg/cerror"
)

const (
	ErrorValidator = "ERROR_VALIDATOR"
)

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DefaultResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ParseToErrorMsg(g *gin.Context, httpStatusCode int, err error) {
	g.AbortWithStatusJSON(httpStatusCode, ErrorResponse{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	})
}

func ParseToDefaultMessage(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, DefaultResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func ParseResponse(ctx *gin.Context, message string, data interface{}, err error) {
	var customErr *cerror.CustomError

	if err != nil {
		if errors.As(err, &customErr) {
			ParseToErrorMsg(ctx, customErr.StatusCode, customErr.Err)
			return
		}

		ParseToErrorMsg(ctx, http.StatusInternalServerError, err)
		return
	}

	ParseToDefaultMessage(ctx, message, data)
}
