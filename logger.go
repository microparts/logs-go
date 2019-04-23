package logs

import (
	"os"
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

//NewLogger is logrus instantiating wrapper. Returns configured logrus instance
func NewLogger(cfg *Config) (*Logger, error) {
	log := &Logger{logrus.New()}
	switch cfg.Format {
	case "text":
		// logg as JSON instead of the default ASCII formatter.
		log.Formatter = &logrus.TextFormatter{
			TimestampFormat:        time.RFC3339,
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			QuoteEmptyFields:       true,
		}
	default:
		// logg as JSON instead of the default ASCII formatter.
		log.Formatter = &logrus.JSONFormatter{TimestampFormat: time.RFC3339}
	}

	log.Level = getLoggerLevel(cfg)

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.Out = os.Stdout

	if cfg.Sentry != nil && cfg.Sentry.Enable {
		err := log.addSentryHook(cfg)
		if err != nil {
			return nil, err
		}
	}

	return log, nil
}

//getLoggerLevel translates text Logger level to logrus Logger level
func getLoggerLevel(cfg *Config) logrus.Level {
	ll, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		ll = logrus.WarnLevel
	}

	return ll
}

func (l *Logger) addSentryHook(cfg *Config) error {
	minLogLevel := getLoggerLevel(cfg)

	hook, err := logrus_sentry.NewSentryHook(cfg.Sentry.DSN, sentryLevels[:minLogLevel])
	if err != nil {
		return err
	}

	hook.SetEnvironment(getSTAGE(cfg))
	hook.Timeout = cfg.Sentry.ResponseTimeout
	if cfg.Sentry.StackTrace.Enable {
		hook.StacktraceConfiguration.Enable = cfg.Sentry.StackTrace.Enable
		hook.StacktraceConfiguration.Level = minLogLevel
		hook.StacktraceConfiguration.Skip = 6
		hook.StacktraceConfiguration.Context = cfg.Sentry.StackTrace.Context
		hook.StacktraceConfiguration.IncludeErrorBreadcrumb = true
		hook.StacktraceConfiguration.SendExceptionType = true
	}

	l.Hooks.Add(hook)

	return nil
}

func getSTAGE(cfg *Config) string {
	if cfg.Sentry.Stage == "" {
		return defaultSTAGE
	}

	return cfg.Sentry.Stage
}
