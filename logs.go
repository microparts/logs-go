package logs

import (
	"github.com/sirupsen/logrus"
)

const (
	ser = "service"
	sta = "stage"
	lat = "latency"

	staI  = "init"
	staS  = "start"
	staDB = "database"
	staQ  = "query"
	staW  = "webserver"
)

var (
	Log = logrus.New()

	// Mux logs
	RouterLogs     = Log.WithField(ser, "router")
	RouterInitLogs = RouterLogs.WithField(sta, staI)
	HttpLogs       = Log.WithField(ser, staW)
	HttpInitLogs   = HttpLogs.WithField(sta, staI)
	HttpStopLogs   = HttpLogs.WithField(sta, "shutdown")

	// Seeds logs
	SeedsLogs = Log.WithField(ser, "seeds").WithField(sta, "seeding")

	// Database logs
	DBLogs     = Log.WithField(ser, staDB)
	DBInitLogs = DBLogs.WithField(sta, staI)

	// CQRS logs
	CommandsLog     = Log.WithField(ser, "CQRS Commands")
	CommandsInitLog = CommandsLog.WithField(sta, staI)
	QueriesLog      = Log.WithField(ser, "CQRS Queries")
	QueriesInitLog  = QueriesLog.WithField(sta, staI)

	// Repository logs
	RepoLog     = Log.WithField(ser, "repository")
	RepoInitLog = RepoLog.WithField(sta, staI)
)
