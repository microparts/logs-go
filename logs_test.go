package logs

import (
	"github.com/sirupsen/logrus"
	"testing"
)

var (
	cfg = &Config{
		LogLevel:  "info",
		LogFormat: "json",
		DSN:       "some_test_string",
	}
	logger *logrus.Logger
)

// I just want to test it before pushing. Don't know how to test it in right way, sorry )
func TestNewLogger(t *testing.T) {
	Init(cfg)
	t.Run("Construct new logger", func(t *testing.T) {
		logger := NewLogger(cfg)
		logger.Debug("some debug")
		logger.Info("some info")
		logger.Warn("some warn")
		logger.Error("some error")
	})
}

func BenchmarkNewLogger(b *testing.B) {
	var l *logrus.Logger
	for n := 0; n < b.N; n++ {
		l = NewLogger(cfg)
	}
	logger = l
}
