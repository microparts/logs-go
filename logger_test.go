package logs

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func logConfig(withSenrty bool) *Config {
	s := &SentryConfig{
		Enable:          true,
		Stage:           "test",
		DSN:             "https://xxx@sentry.io/yyy",
		ResponseTimeout: 0,
		MinlLogLevel:    "error",
		StackTrace: StackTraceConfig{
			Enable: true,
		},
	}
	cfg := &Config{
		Level:  "info",
		Format: "text",
	}
	if withSenrty {
		cfg.Sentry = s
	}
	return cfg
}

// I just want to test it before pushing. Don't know how to test it in right way, sorry )
func TestNewLogger(t *testing.T) {
	t.Run("Construct new logger", func(t *testing.T) {
		cfg := logConfig(false)
		l, err := NewLogger(cfg)
		assert.NoError(t, err)
		assert.Equal(t, logrus.InfoLevel, l.GetLevel())
	})

	t.Run("Logger format", func(t *testing.T) {
		cfg := logConfig(false)
		cfg.Format = "du'soley"
		l, err := NewLogger(cfg)
		assert.NoError(t, err)
		assert.Equal(t, &logrus.JSONFormatter{TimestampFormat: time.RFC3339}, l.Formatter)
	})

	t.Run("good sentry config", func(t *testing.T) {
		cfg := logConfig(true)
		l, err := NewLogger(cfg)
		assert.NoError(t, err)
		assert.Equal(t, len(l.Hooks), len(sentryLevels[:getLoggerLevel(cfg)]))
	})

	t.Run("bad sentry config", func(t *testing.T) {
		cfg := logConfig(true)
		cfg.Sentry.DSN = "bad dsn"
		_, err := NewLogger(cfg)
		assert.Error(t, err)
	})

}

func TestGetLoggerLeve(t *testing.T) {
	t.Run("existing Logger level", func(t *testing.T) {
		cfg := logConfig(false)
		ll := getLoggerLevel(cfg)
		assert.Equal(t, logrus.InfoLevel, ll)

	})
	t.Run("not existing Logger level", func(t *testing.T) {
		cfg := logConfig(false)
		cfg.Level = "paranoia"
		ll := getLoggerLevel(cfg)
		assert.Equal(t, logrus.WarnLevel, ll)
	})
}

func TestSentryStage(t *testing.T) {
	t.Run("stage is set", func(t *testing.T) {
		cfg := logConfig(true)
		assert.Equal(t, "test", getSTAGE(cfg))
	})

	t.Run("stage is not set", func(t *testing.T) {
		cfg := logConfig(true)
		cfg.Sentry.Stage = ""
		assert.Equal(t, defaultSTAGE, getSTAGE(cfg))
	})
}
