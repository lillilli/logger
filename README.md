# Logger

[![Go Report Card](https://goreportcard.com/badge/lillilli/jsonrpc)](https://goreportcard.com/report/lillilli/logger)
[![GoDoc](https://godoc.org/github.com/lillilli/jsonrpc?status.svg)](https://godoc.org/github.com/lillilli/logger)
[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/lillilli/logger/master/LICENSE)

Simple logger broker under hashicorp/logutils.

## Description

Simple logger, that implement specified interface (logger.Logger). That logger can log into anything, that implemented io.Writer interface.

## Base

Logger params setup by one call of Init(logger.Prams) method:

```go
// Params - logger params
type Params struct {
    // Output interface
    Writer   io.Writer
    // Log levels
    Levels   []string
    // Min log level (all logs, that stand before that level will not be logged)
    MinLevel string
}
```

Logger interface:

```go
// Logger - logger interface
type Logger interface {
    Debug(msg string)
    Debugf(msg string, args ...interface{})

    Info(msg string)
    Infof(format string, args ...interface{})

    Warn(msg string)
    Warnf(format string, args ...interface{})

    Error(msg string)
    Errorf(format string, args ...interface{})

    Fatal(msg string)
    Fatalf(format string, args ...interface{})
}
```

### Log format

```bash
# <data> <time> <severity> <module>: <message>

2018/05/08 09:28:49 [INFO] service: Starting...
```

## Usage examples

### stderr

```go
package main

import (
	"github.com/lillilli/logger"
)

type Service struct {
	log logger.Logger
}

func NewService() *Service {
	return &Service{
		log: logger.NewLogger("service name"),
	}
}

func (s *Service) SayHi(name string) {
	s.log.Infof("Saying hi to %s", name)
}

func main() {
	service := NewService()
	service.SayHi("Alex")

	// Output:
	// 2019/03/12 11:56:53 [INFO] service name: Saying hi to Alex
}
```

## GELF (graylog)

```go
import "gopkg.in/Graylog2/go-gelf.v1/gelf"

gelfWriter, err := gelf.NewWriter("localhost:12201")
if err != nil {
    return errors.Wrap(err, "unable to create gelf writer")
}
```

### syslog

```go
import (
  "log"
	"log/syslog"

	"github.com/lillilli/logger"
)

func main() {
  logWriter, err := syslog.New(syslog.LOG_NOTICE, "service_name")

	if err != nil {
		log.Fatalf("Unable to create syslog writer: %v", err)
	}

	logger.Init(logger.Params{
		Writer: logWriter,
	})

	log := logger.NewLogger("service")
	log.Info("I'm going to syslog")
}
```

### rsyslog (udp)

```go
logWriter, err := syslog.Dial("udp", "rsyslog:514", syslog.LOG_NOTICE, "service_name")

if err != nil {
    return errors.Wrap(err, "unable to create syslog writer")
}
```

### rsyslog (tcp)

```go
logWriter, err := syslog.Dial("tcp", "rsyslog:10514", syslog.LOG_NOTICE, "service_name")

if err != nil {
    return errors.Wrap(err, "unable to create syslog writer")
}
```

## License

Released under the [MIT License](https://github.com/lillilli/logger/blob/master/LICENSE).