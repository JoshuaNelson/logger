package logger

import (
	"log"
	"os"
)

var LOG_FILE = "/home/joshuanelsn/go/src/bsh-tfe/debug.log"
var logFile *os.File

var debug *log.Logger

func Init() {
	f, err := os.OpenFile(LOG_FILE, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logFile = f
}

func Debug(format string, v ...interface{}) {
	if debug == nil {
		debug = log.New(logFile, "DEBUG: ", log.Ltime)
	}
	debug.Printf(format, v...)
}

func Close() {
	logFile.Close()
}
