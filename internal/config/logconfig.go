package config

import (
	"github.com/sirupsen/logrus"
)

const (
	defaultLogLevel = logrus.InfoLevel
)

func LogConfig() {
	if logLevel, err := logrus.ParseLevel(cfg.LogLevel); err == nil {
		logrus.SetLevel(logLevel)
	} else {
		logrus.SetLevel(defaultLogLevel)
		logrus.Error(err)
		logrus.Warn("log level set by default value")
	}

	// f, err := os.OpenFile("filename", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	// if err != nil {
	// 	logrus.Error("error of openning a file for logging: ", err)
	// }
	// logrus.SetOutput(f)
}
