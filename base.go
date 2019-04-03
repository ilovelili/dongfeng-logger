package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

// Log log entity
type Log struct {
	Category string
	Content  string
	Time     time.Time
}

// Level logger level
type Level int

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

// Option logger option
type Option struct {
	Level  Level
	NoLock bool
}

// DefaultOption default logger option
func DefaultOption() *Option {
	return &Option{
		Level:  InfoLevel,
		NoLock: false,
	}
}

var (
	operationlogger *logrus.Logger
	systemlogger    *logrus.Logger
	errorlogger     *logrus.Logger
)

// OperationLog opertion level log
func (l *Log) OperationLog(index Index) {
	option := &Option{Level: InfoLevel}
	operationlogger = StdOutLogger(index, option)
	operationlogger.Info(l)
}

// SystemLog system level log
func (l *Log) SystemLog(index Index) {
	option := &Option{Level: InfoLevel}
	systemlogger = StdOutLogger(index, option)
	systemlogger.Info(l)
}

// ErrorLog error level log
func (l *Log) ErrorLog(index Index) {
	option := &Option{Level: ErrorLevel}
	errorlogger = StdOutLogger(index, option)
	errorlogger.Error(l)
}

// Index logger index
type Index string

// extend logger index here
const (
	// CoreServer core server
	CoreServer Index = "core_server"
	// CoreProxy core proxy
	CoreProxy Index = "core_proxy"
	// WebSocket websocket
	WebSocket Index = "websocket"
	// PhysiqueServer physique
	PhysiqueServer Index = "physique_server"
	// NutritionServer nutrition
	NutritionServer Index = "nutrition_server"
	// CronJobs cron jobs
	CronJobs Index = "cron_jobs"
)
