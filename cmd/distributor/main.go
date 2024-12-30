package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	configureLog()
}

func configureLog() {
	// TODO: logrus.SetLevel(logrus.InfoLevel)
	logrus.SetLevel(logrus.DebugLevel)

	logrus.SetReportCaller(false)

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006.01.02 15:04:05.000",
		ForceColors:     true,
	})

	logrus.SetOutput(os.Stdout)
}

func measureTime(action func()) (duration time.Duration) {
	start := time.Now()
	action()
	duration = time.Since(start)

	return
}
