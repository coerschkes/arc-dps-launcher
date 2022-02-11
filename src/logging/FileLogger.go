package logging

import (
	"log"
	"os"
)

/*
	Author: Christian Oerschkes <christian.oerschkes@hotmail.de>
*/

type FileLogger struct {
	logFile  string
	logLevel LOG_LEVEL
}

const DEFAULT_LOG_FILE = "Log.log"

var loggerInstance Logger

func init() {
	loggerInstance = &FileLogger{DEFAULT_LOG_FILE, INFO}
}

func GetLogger() Logger {
	return loggerInstance
}

func (l *FileLogger) Log(message string) {
	l.LogWithLogLevel(message, l.logLevel)
}

func (l *FileLogger) LogWithLogLevel(message string, logLevel LOG_LEVEL) {
	if logFile, err := os.OpenFile(l.logFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644); err != nil {
		log.Panic(err)
	} else {
		defer logFile.Close()
		log.SetOutput(logFile)
		log.SetFlags(log.LstdFlags)
		log.Println(string(l.logLevel) + ": " + message)
	}
}

func (l *FileLogger) SetDefaultLogLevel(level LOG_LEVEL) {
	l.logLevel = level
}

func (l *FileLogger) SetOutputFile(path string) {
	l.logFile = path
}
