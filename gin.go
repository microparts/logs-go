package logs

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

// MuxLogger Логирование работы веб-сервера
func MuxLogger() gin.HandlerFunc {
	var skip map[string]struct{}
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		if _, ok := skip[path]; !ok {
			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()

			comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

			if raw != "" {
				path = path + "?" + raw
			}

			// Stop timer
			end := time.Now()
			latency := end.Sub(start)

			Log.WithFields(logrus.Fields{
				ser:        staW,
				sta:        "request",
				"latency":  latency,
				"clientIP": clientIP,
				"status":   statusCode,
				"proto":    c.Request.Proto,
				"method":   method,
				"path":     path,
				staQ:       raw,
				"comment":  comment,
			}).Info("Incoming request")
		}
	}
}
