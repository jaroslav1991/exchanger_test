package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

func GetHostConfig() string {
	if host := os.Getenv("HOST"); host != "" {
		return host
	}

	return "localhost"
}

func GetPortConfig() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}

	return ":8080"
}

func GetLoggerLevel() logrus.Level {
	level := os.Getenv("LOG_LEVEL")
	if level != "" {
		switch level {
		case "panic":
			return logrus.PanicLevel
		case "fatal":
			return logrus.FatalLevel
		case "error":
			return logrus.ErrorLevel
		case "warn":
			return logrus.WarnLevel
		case "info":
			return logrus.InfoLevel
		case "debug":
			return logrus.DebugLevel
		case "trace":
			return logrus.TraceLevel
		}
	}

	return logrus.DebugLevel
}
