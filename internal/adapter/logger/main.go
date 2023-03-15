package logger

import (
	"github.com/BasbustDama/note-service/internal/core/interfaces"
	"github.com/sirupsen/logrus"
)

type logger struct {
	log *logrus.Logger
}

func New(log *logrus.Logger) interfaces.Logger {
	return &logger{
		log: log,
	}
}

func (l *logger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}
