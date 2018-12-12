package gocar

import (
	"io"
	"net/http"
	"time"
)

func errorHandler(err error, res http.ResponseWriter, info ...interface{}) bool {

	if err != nil {
		warnLog(err, info)
		io.WriteString(res, "404 Page Not Found")
		return true
	}
	return false
}

func HandlePath(path string, handler func(http.ResponseWriter, *http.Request) error, middlewares ...func(http.ResponseWriter, *http.Request) bool) {
	mux.HandleFunc(path, func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		en := l.isLoggerEnabled()

		if en == true {
			l.Initialize(res, req)
		}

		var err error

		func() {
			for _, f := range middlewares {
				next := f(res, req)

				if !next {
					return
				}
			}
			err = handler(res, req)
		}()

		elapsed := time.Since(startTime).String()

		if enableLogger == true {
			if !en && !errorHandler(err, res, req.Method, path) {
				infoLog(req.Method, path, elapsed)
			} else if en == true {
				l.Log(res, req)
			}
		}
	})

	autoSetHandler()
}
