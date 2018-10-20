package logger

import (
	"os"

	"github.com/Sirupsen/logrus"
)

// StdOutLogger init a stdout logger
func StdOutLogger(index Index, option ...*Option) *logrus.Logger {
	var Option *Option
	if len(option) == 0 {
		Option = DefaultOption()
	} else {
		Option = option[0]
	}

	log := logrus.New()
	log.Out = os.Stdout
	log.SetLevel(logrus.Level(Option.Level))
	log.Formatter = new(logrus.JSONFormatter)
	log.WithField("index", index)

	if Option.NoLock {
		log.SetNoLock()
	}

	return log
}
