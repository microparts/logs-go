package logs

import (
	"github.com/gin-gonic/gin"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type DBLogger struct{}

type Config struct {
	LogLevel   string
	LogFormat  string
	TimeFormat string
}

const (
	ser = "service"
	sta = "stage"

	staI  = "init"
	staDB = "database"
	staQ  = "query"
	staW  = "webserver"
)

var (
	Log = logrus.New()

	// Mux logs
	RouterLogs     = Log.WithField(ser, "router")
	RouterInitLogs = RouterLogs.WithField(sta, staI)
	HttpLogs       = Log.WithField(ser, staW)
	HttpInitLogs   = HttpLogs.WithField(sta, staI)
	HttpStopLogs   = HttpLogs.WithField(sta, "shutdown")

	// Bot logs
	BotLogs     = Log.WithField(ser, "bot")
	BotInitLogs = BotLogs.WithField(sta, staI)

	// Seeds logs
	SeedsLogs = Log.WithField(ser, "seeds").WithField(sta, "seeding")

	// Database logs
	DBLogs     = Log.WithField(ser, staDB)
	DBInitLogs = DBLogs.WithField(sta, staI)

	// CQRS logs
	CommandsLog     = Log.WithField(ser, "CQRS Commands")
	CommandsInitLog = CommandsLog.WithField(sta, staI)
	QueriesLog      = Log.WithField(ser, "CQRS Queries")
	QueriesInitLog  = QueriesLog.WithField(sta, staI)

	// Repository logs
	RepoLog     = Log.WithField(ser, "repository")
	RepoInitLog = RepoLog.WithField(sta, staI)
)

// LogsInit Инициация логгера
func LogsInit(logCfg *Config) {

	switch logCfg.LogFormat {
	case "text":
		// logg as JSON instead of the default ASCII formatter.
		Log.Formatter = &logrus.TextFormatter{
			TimestampFormat:        logCfg.TimeFormat,
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			QuoteEmptyFields:       true,
		}
	default:
		// logg as JSON instead of the default ASCII formatter.
		Log.Formatter = &logrus.JSONFormatter{TimestampFormat: logCfg.TimeFormat}
	}

	switch logCfg.LogLevel {
	case "panic":
		Log.Level = logrus.PanicLevel
	case "fatal":
		Log.Level = logrus.FatalLevel
	case "error":
		Log.Level = logrus.ErrorLevel
	case "warn":
		Log.Level = logrus.WarnLevel
	case "info":
		Log.Level = logrus.InfoLevel
	case "debug":
		Log.Level = logrus.DebugLevel
	default:
		Log.Level = logrus.WarnLevel
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	Log.Out = os.Stdout

	Log.WithFields(logrus.Fields{ser: "log", sta: staI}).Info("Logs initiated")
}

// Лоирование работы БД
func (*DBLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		Log.WithFields(logrus.Fields{ser: staDB, sta: staQ, "sql": v[3], "values": v[4]}).Info("Query sql")
	}
	if v[0] == "log" {
		Log.WithFields(logrus.Fields{ser: staDB, sta: staQ, staQ: v[2]}).Info("Query log")
	}
}

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
