package gin

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MuxLogger struct {
	l *logrus.Logger
}

func NewLogger(logger *logrus.Logger) *MuxLogger {
	return &MuxLogger{l: logger}
}

// MuxLogger Логирование работы веб-сервера
func (m *MuxLogger) Log() gin.HandlerFunc {
	var skip map[string]struct{}
	// @TODO попробовать избавиться от зависимости на gin либу
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

			m.l.WithFields(logrus.Fields{
				"latency":  latency,
				"clientIP": clientIP,
				"status":   statusCode,
				"proto":    c.Request.Proto,
				"method":   method,
				"path":     path,
				"query":    raw,
				"comment":  comment,
			}).Info("http request")
		}
	}
}
