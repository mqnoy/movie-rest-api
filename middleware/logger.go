package middleware

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mqnoy/movie-rest-api/pkg/logger"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var buf bytes.Buffer
		tee := io.TeeReader(ctx.Request.Body, &buf)
		body, err := io.ReadAll(tee)
		ctx.Request.Body = io.NopCloser(&buf)

		// next to processing request
		ctx.Next()

		reqMethod := ctx.Request.Method
		reqUri := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()

		if statusCode != http.StatusOK {
			logger.Error().
				Err(err).
				Str("context", "http.call").
				Str("method", reqMethod).
				Str("endpoint", reqUri).
				Int("status", statusCode).
				Str("data", string(body)).
				Send()
		}

		ctx.Next()
	}
}
