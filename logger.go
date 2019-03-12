package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/logutils"
)

// Simple broker under logutils

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

// Params - logger params
type Params struct {
	Writer   io.Writer
	Levels   []string
	MinLevel string `env:"LOG_LEVEL" default:"DEBUG"`
}

type logger struct {
	packageName string
}

// Init - logger initialization
func Init(params Params) {
	levels := []logutils.LogLevel{"DEBUG", "WARN", "INFO", "ERROR", "FATAL"}
	minLevel := logutils.LogLevel("DEBUG")

	var writer io.Writer = os.Stderr

	// Setuping different log levels (by default ["DEBUG", "WARN", "INFO", "ERROR", "FATAL"])
	if params.Levels != nil {
		levels = []logutils.LogLevel{}

		for _, level := range params.Levels {
			levels = append(levels, logutils.LogLevel(strings.ToUpper(level)))
		}
	}

	// Setuping min logging level (DEBUG by default)
	if params.MinLevel != "" {
		minLevel = logutils.LogLevel(strings.ToUpper(params.MinLevel))
	}

	// Setuping output writer for log (os.Stderr by default)
	if params.Writer != nil {
		writer = params.Writer
	}

	filter := &logutils.LevelFilter{
		Levels:   levels,
		MinLevel: minLevel,
		Writer:   writer,
	}

	log.SetOutput(filter)
}

// NewLogger - return new logger instance
func NewLogger(packageName string) Logger {
	return logger{packageName: packageName + ":"}
}

func (l logger) Debug(msg string) {
	log.Println("[DEBUG]", l.packageName, msg)
}

func (l logger) Debugf(format string, args ...interface{}) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l logger) Info(msg string) {
	log.Println("[INFO]", l.packageName, msg)
}

func (l logger) Infof(format string, args ...interface{}) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l logger) Warn(msg string) {
	log.Println("[WARN]", l.packageName, msg)
}

func (l logger) Warnf(format string, args ...interface{}) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l logger) Error(msg string) {
	log.Println("[ERROR]", l.packageName, msg)
}

func (l logger) Errorf(format string, args ...interface{}) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l logger) Fatal(msg string) {
	log.Println("[FATAL]", l.packageName, msg)
}

func (l logger) Fatalf(format string, args ...interface{}) {
	l.Fatal(fmt.Sprintf(format, args...))
}
