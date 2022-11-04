package logging

import (
	"log"
	"os"
)

type logger struct {
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
}

var Log *logger = nil

func Init() {

	infoLogger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	Log = &logger{warningLogger: warningLogger, infoLogger: infoLogger, errorLogger: errorLogger}
}

func (l *logger) Error(message string) {
	l.errorLogger.Println(message)
}

func (l *logger) Info(message string) {
	l.infoLogger.Println(message)
}

func (l *logger) Warn(message string) {
	l.warningLogger.Println(message)
}
