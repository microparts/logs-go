package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	environmentVar     = "STAGE"
	defaultEnvironment = "unknown"
)

type Config struct {
	LogLevel   string `yaml:"level"`
	LogFormat  string `yaml:"format"`
	TimeFormat string
	DSN        string `yaml:"dsn"`
	StackTrace struct {
		Skip    int `yaml:"skip"`
		Context int `yaml:"context"`
	} `yaml:"stacktrace"`
}

// LogsInit Инициация логгера
func Init(logCfg *Config) {
	configs = logCfg
	switch configs.LogFormat {
	case "text":
		// logg as JSON instead of the default ASCII formatter.
		Log.Formatter = &logrus.TextFormatter{
			TimestampFormat:        configs.TimeFormat,
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			QuoteEmptyFields:       true,
		}
	default:
		// logg as JSON instead of the default ASCII formatter.
		Log.Formatter = &logrus.JSONFormatter{TimestampFormat: configs.TimeFormat}
	}

	switch configs.LogLevel {
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

	if configs.DSN != "" {
		hook, err := NewSentryHook(configs.DSN)
		if err == nil {
			env := os.Getenv(environmentVar)
			if env == "" {
				env = defaultEnvironment
			}
			hook.client.SetEnvironment(env)

			Log.Hooks.Add(hook)
		}
	}

	Log.WithFields(logrus.Fields{ser: "log", sta: staI}).Info("Logs initiated")
}
