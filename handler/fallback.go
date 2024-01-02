package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FallbackHandler(ctx *gin.Context) {
	message := fmt.Sprintf("cant find %s", ctx.Request.URL.Path)
	ctx.AbortWithStatusJSON(http.StatusNotFound, map[string]any{
		"message": message,
	})
}
