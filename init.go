package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	LogLevel   string `yaml:"level"`
	LogFormat  string `yaml:"format"`
	TimeFormat string
	DSN        string `yaml:"dsn"`
}

// LogsInit Инициация логгера
func Init(logCfg *Config) {

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

	if logCfg.DSN != "" {
		hook, err := NewSentryHook(logCfg.DSN)
		if err == nil {
			Log.Hooks.Add(hook)
		}
	}

	Log.WithFields(logrus.Fields{ser: "log", sta: staI}).Info("Logs initiated")
}
