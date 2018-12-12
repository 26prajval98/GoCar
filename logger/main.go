package logger

import (
	"log"
	"os"
)

var (
	Info *log.Logger
	Warn *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "INFO :", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stderr, "ERR :", log.Ldate|log.Ltime|log.Lshortfile)
}

func InfoLog(y ...interface{}) {
	Info.Println(y)
}

func WarnLog(y ...interface{}) {
	Warn.Println(y)
}
