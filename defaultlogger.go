package gocar

import (
	"log"
	"os"
)

var (
	info *log.Logger
	warn *log.Logger
)

func init() {
	info = log.New(os.Stdout, "INFO :", log.Ldate|log.Ltime|log.Lshortfile)
	warn = log.New(os.Stderr, "ERR :", log.Ldate|log.Ltime|log.Lshortfile)
}

func infoLog(y ...interface{}) {
	info.Println(y)
}

func warnLog(y ...interface{}) {
	warn.Println(y)
}
