package logs

import (
	"time"

	"github.com/sirupsen/logrus"
)

const (
	defaultSTAGE = "defaults"
)

type StackTraceConfig struct {
	Enable  bool `yaml:"enable"`
	Context int  `yaml:"context"`
}

type SentryConfig struct {
	Enable          bool             `yaml:"enable"`
	Stage           string           `yaml:"stage"`
	MinlLogLevel    string           `yaml:"min_log_level"`
	DSN             string           `yaml:"dsn"`
	ResponseTimeout time.Duration    `yaml:"response_timeout"`
	StackTrace      StackTraceConfig `yaml:"stacktrace"`
}

type Config struct {
	Level  string        `yaml:"level"`
	Format string        `yaml:"format"`
	Sentry *SentryConfig `yaml:"sentry,omitempty"`
}

var (
	sentryLevels = []logrus.Level{
		logrus.DebugLevel,
		logrus.InfoLevel,
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
)
