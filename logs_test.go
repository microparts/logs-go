package logs

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func logConfig() *Config {
	return &Config{
		Level:  "info",
		Format: "text",
	}
}

// I just want to test it before pushing. Don't know how to test it in right way, sorry )
func TestNewLogger(t *testing.T) {
	cfg := logConfig()

	t.Run("Construct new logger", func(t *testing.T) {
		l, err := NewLogger(cfg)
		assert.NoError(t, err)
		assert.Equal(t, logrus.InfoLevel, l.Level)
	})

	if cfg.Sentry != nil && cfg.Sentry.Enable {
		t.Run("test sentry hook", func(t *testing.T) {
			l, err := NewLogger(cfg)
			assert.NoError(t, err)
			l.Error("some error fire to sentry")
			l.Warning("some warning fire to sentry")
		})
	}

}
