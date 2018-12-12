package gocar

import (
	"log"
	"net/http"
	"strconv"
)

type Logger struct {
	Initialize func(http.ResponseWriter, *http.Request)
	Log        func(http.ResponseWriter, *http.Request)
}

var l Logger

var mux = http.NewServeMux()

var handler http.Handler

var enableLogger bool

func init() {
	autoSetHandler()
	SetLoggerState(true)
}

func (l Logger) isLoggerEnabled() bool {
	if l.Initialize == nil && l.Log == nil {
		return false
	}
	return true
}

func autoSetHandler() {
	handler = mux
}

func SetHandler(h func(http.Handler) http.Handler) {
	handler = h(mux)
}

func SetLoggerState(flag bool) {
	enableLogger = flag
}

func SetLogger(logPack Logger) {
	l = logPack
}

func StartServer(port int) {
	p := strconv.Itoa(port)
	p = ":" + p
	log.Fatal(http.ListenAndServe(p, handler))
}
