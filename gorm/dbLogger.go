package gorm

import (
	"github.com/sirupsen/logrus"
)

type DBLogger struct {
	l *logrus.Logger
}

func NewLogger(logger *logrus.Logger) *DBLogger {
	return &DBLogger{l: logger}
}

// Лоирование работы БД
func (d *DBLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		d.l.WithFields(logrus.Fields{"query": v[3], "values": v[4], "rows": v[5], "latency": v[2]}).Info("Query sql")
	}
	if v[0] == "log" {
		d.l.WithFields(logrus.Fields{"query": v[2]}).Info("Query log")
	}
}
