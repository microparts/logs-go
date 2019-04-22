package logs

import (
	"os"
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/sirupsen/logrus"
)

const (
	defaultEnvironment = "unknown"
)

type StackTraceConfig struct {
	Enable  bool
	Context int `yaml:"context"`
}

type SentryConfig struct {
	Enable          bool             `yaml:"enable"`
	Stage           string           `yaml:"stage"`
	DSN             string           `yaml:"dsn"`
	ResponseTimeout time.Duration    `yaml:"response_timeout"`
	StackTrace      StackTraceConfig `yaml:"stacktrace"`
}

type Config struct {
	Level  string        `yaml:"level"`
	Format string        `yaml:"format"`
	Sentry *SentryConfig `yaml:"sentry,omitempty"`
}

type Log struct {
	l *logrus.Logger
}

// Logger constructor
// Returns new logger instance
func NewLogger(cfg *Config) (*logrus.Logger, error) {
	log := &Log{l: logrus.New()}
	switch cfg.Format {
	case "text":
		// logg as JSON instead of the default ASCII formatter.
		log.l.Formatter = &logrus.TextFormatter{
			TimestampFormat:        time.RFC3339,
			FullTimestamp:          true,
			DisableLevelTruncation: true,
			QuoteEmptyFields:       true,
		}
	default:
		// logg as JSON instead of the default ASCII formatter.
		log.l.Formatter = &logrus.JSONFormatter{TimestampFormat: time.RFC3339}
	}

	var sentryLevels []logrus.Level
	switch cfg.Level {
	case "panic":
		log.l.Level = logrus.PanicLevel
		sentryLevels = []logrus.Level{logrus.PanicLevel}
	case "fatal":
		log.l.Level = logrus.FatalLevel
		sentryLevels = []logrus.Level{logrus.PanicLevel, logrus.FatalLevel}
	case "error":
		log.l.Level = logrus.ErrorLevel
		sentryLevels = []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel}
	case "warn":
		log.l.Level = logrus.WarnLevel
		sentryLevels = []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.WarnLevel}
	case "info":
		log.l.Level = logrus.InfoLevel
		sentryLevels = []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.WarnLevel}
	case "debug":
		log.l.Level = logrus.DebugLevel
		sentryLevels = []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.WarnLevel}
	default:
		log.l.Level = logrus.WarnLevel
		sentryLevels = []logrus.Level{logrus.PanicLevel, logrus.FatalLevel, logrus.ErrorLevel, logrus.WarnLevel}
	}

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.l.Out = os.Stdout

	if cfg.Sentry != nil && cfg.Sentry.Enable {
		err := log.addSentryHook(cfg, sentryLevels)
		if err != nil {
			return nil, err
		}
	}

	return log.l, nil
}

func (log *Log) addSentryHook(cfg *Config, logLevels []logrus.Level) error {
	hook, err := logrus_sentry.NewSentryHook(cfg.Sentry.DSN, logLevels)
	if err != nil {
		return err
	}

	if cfg.Sentry.Stage == "" {
		cfg.Sentry.Stage = defaultEnvironment
	}

	hook.SetEnvironment(cfg.Sentry.Stage)
	hook.Timeout = cfg.Sentry.ResponseTimeout
	if cfg.Sentry.StackTrace.Enable {
		hook.StacktraceConfiguration.Enable = cfg.Sentry.StackTrace.Enable
		hook.StacktraceConfiguration.Level = logrus.ErrorLevel
		hook.StacktraceConfiguration.Skip = 6
		hook.StacktraceConfiguration.Context = cfg.Sentry.StackTrace.Context
		hook.StacktraceConfiguration.IncludeErrorBreadcrumb = true
		hook.StacktraceConfiguration.SendExceptionType = true
	}

	log.l.Hooks.Add(hook)

	return nil
}
