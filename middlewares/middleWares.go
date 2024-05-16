package middlewares

import (
	"fmt"
	"net/http"
	"practice/server/router"
)

func LoggingMiddleWare() router.MiddleWareFunc {
	return func(w http.ResponseWriter, r *http.Request, next func()) {
		fmt.Print("a new request received\n")
		next()
	}
}

func ContentTypeMiddleWare() router.MiddleWareFunc {
	return func(w http.ResponseWriter, r *http.Request, next func()) {
		w.Header().Set("Content-Type", "application-json")
		next()
	}
}
