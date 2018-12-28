package logs

import (
	"github.com/sirupsen/logrus"
)

type DBLogger struct{}

// Лоирование работы БД
func (*DBLogger) Print(v ...interface{}) {
	if v[0] == "sql" {
		Log.WithFields(logrus.Fields{ser: staDB, sta: staQ, "query": v[3], "values": v[4], "rows": v[5], lat: v[2]}).Info("Query sql")
	}
	if v[0] == "log" {
		Log.WithFields(logrus.Fields{ser: staDB, sta: staQ, staQ: v[2]}).Info("Query log")
	}
}
