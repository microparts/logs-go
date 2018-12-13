package logs

import (
	"github.com/sirupsen/logrus"
	"os"
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
	staS  = "start"
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
	BotLogs      = Log.WithField(ser, "bot")
	BotInitLogs  = BotLogs.WithField(sta, staI)
	BotStartLogs = BotLogs.WithField(sta, staS)

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
