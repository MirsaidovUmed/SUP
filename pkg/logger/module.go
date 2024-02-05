package logger

import (
	"SUP/pkg/config"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(config *config.Config) *logrus.Logger {
	f, err := os.OpenFile(config.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o777)
	if err != nil {
		panic(err)
	}

	log := &logrus.Logger{
		// Log into f file handler and on os.Stdout
		Out:       io.MultiWriter(f, os.Stdout),
		Level:     logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{},
	}
	log.SetReportCaller(true)

	return log
}
