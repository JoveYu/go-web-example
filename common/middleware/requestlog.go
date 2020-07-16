package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func RequestLog() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		end := time.Now()

		latency := end.Sub(start)
		msg := "gin"
		if len(ctx.Errors) > 0 {
			msg = ctx.Errors.String()
		}
		status := ctx.Writer.Status()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery
		if query != "" {
			path = fmt.Sprintf("%s?%s", path, query)
		}

		logger := log.With().
			Int("status", status).
			Str("method", ctx.Request.Method).
			Str("path", path).
			Str("ip", ctx.ClientIP()).
			Int64("duration", int64(latency/time.Millisecond)). // ms
			Str("user-agent", ctx.Request.UserAgent()).
			Logger()

		if status >= 400 && status < 500 {
			logger.Warn().Msg(msg)
		} else if status > 500 {
			logger.Error().Msg(msg)
		} else {
			logger.Info().Msg(msg)
		}

	}
}
