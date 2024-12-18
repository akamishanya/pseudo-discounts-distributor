package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	configureLog()

	start := time.Now()
	// ...
	duration := time.Since(start)

	logrus.Infoln("The program finished in ", duration)
}

func configureLog() {
	// Установка уровня логирования (TODO: заменить на InfoLevel)
	logrus.SetLevel(logrus.DebugLevel)

	// Установка отображения вызывающей функции
	logrus.SetReportCaller(false)

	// Установка форматтера
	logrus.SetFormatter(&logrus.TextFormatter{
		// Полный формат времени
		FullTimestamp: true,

		// Дата и время в формате "день.месяц.год час:минута:секунда:миллисекунда"
		TimestampFormat: "02.01.2006 15:04:05:000",

		// Принудительное использование цветов
		ForceColors: true,
	})

	// Установка логирования в stdout
	logrus.SetOutput(os.Stdout)
}
