logs-go
-------

[Logrus](github.com/sirupsen/logrus) wrapper for easy use with sentry hook, database (gorm) and mux (gin) loggers.

## Usage

Initiate new logger with prefilled `logs.Config` and use it as common logrus logger instance

```go
package main

import (
	"time"
	
	"github.com/microparts/logs-go"
)

func main() {
	conf := &logs.Config{
		Level:"warn",
		Format: "text",
		Sentry: &logs.SentryConfig{
			Enable: true,
			Stage:"test",
			DSN: "http://dsn.sentry.com",
			ResponseTimeout: 0,
			StackTrace: logs.StackTraceConfig{
				Enable: true,
			},
		},
	}
	
	l, err := logs.NewLogger(conf)
	if err != nil {
		panic(err)
	}
	
	l.Warn("log some warning")
}
```

## Licence

The software is provided under [MIT Licence](LICENCE).