package logging

import (
	"log"
	"os"
)

/*
	Author: Christian Oerschkes <christian.oerschkes@hotmail.de>
*/
type FileLogger struct {
	srcFile string
	logFile string
}

const DEFAULT_LOG_FILE = "Log.log"

func GetLogger(src string) Logger {
	return &FileLogger{src, DEFAULT_LOG_FILE}
}

func (l FileLogger) Log(message string) {
	if logFile, err := os.OpenFile(l.logFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644); err != nil {
		log.Panic(err)
	} else {
		defer logFile.Close()
		log.SetOutput(logFile)
		log.SetFlags(log.LstdFlags)
		log.Println(l.srcFile + ": " + message)
	}
}

func (l *FileLogger) SetOutputFile(path string) {
	l.logFile = path
}
