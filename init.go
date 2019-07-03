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

// Logger constructor
// Returns new logger instance
func NewLogger(cfg *Config) *logrus.Logger {
	configs = cfg
	log := logrus.New()
	switch configs.LogFormat {
	case "text":
		// logg as JSON instead of the default ASCII formatter.
		log.Formatter = &logrus.TextFormatter{
			TimestampFormat:        configs.TimeFormat,
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			QuoteEmptyFields:       true,
		}
	default:
		// logg as JSON instead of the default ASCII formatter.
		log.Formatter = &logrus.JSONFormatter{TimestampFormat: configs.TimeFormat}
	}

	switch configs.LogLevel {
	case "panic":
		log.Level = logrus.PanicLevel
	case "fatal":
		log.Level = logrus.FatalLevel
	case "error":
		log.Level = logrus.ErrorLevel
	case "warn":
		log.Level = logrus.WarnLevel
	case "info":
		log.Level = logrus.InfoLevel
	case "debug":
		log.Level = logrus.DebugLevel
	default:
		log.Level = logrus.WarnLevel
	}

	Log.Level = log.Level

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.Out = os.Stdout

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

	return log
}

// LogsInit Инициация логгера
// Deprecated
func Init(logCfg *Config) {
	Log = NewLogger(logCfg)
	Log.WithFields(logrus.Fields{ser: "log", sta: staI}).Info("Logs initiated")

}
