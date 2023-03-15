package main

import (
	"github.com/BasbustDama/note-service/internal/adapter/logger"
	"github.com/BasbustDama/note-service/pkg/logrus"
)

func main() {
	// packages
	logrus := logrus.New()

	// adapter
	logger := logger.New(logrus)

	// TODO: убрать нижнюю строчку
	_ = logger
}
