package logs

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func logConfig() *Config {
	return &Config{
		Env:    "test",
		Level:  "info",
		Format: "test",
	}
}

// I just want to test it before pushing. Don't know how to test it in right way, sorry )
func TestNewLogger(t *testing.T) {
	cfg := logConfig()
	l := NewLogger(cfg)

	t.Run("Construct new logger", func(t *testing.T) {
		assert.Equal(t, logrus.InfoLevel, l.Level)
	})

	if cfg.Sentry.Enable {
		t.Run("test sentry hook", func(t *testing.T) {
			l.Error("some error fire to sentry")

			l.Warning("some warning fire to sentry")
			//l.Fatal("some fatal error fire to sentry")
		})
	}

}
