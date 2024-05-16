package router

import (
	"net/http"
)

type MiddleWareFunc = func(w http.ResponseWriter, r *http.Request, next func())

type Router struct {
	BaseServer  *http.ServeMux
	middleWares []MiddleWareFunc
}

type middleWareExecutor struct {
	index int
}

func (executor middleWareExecutor) execute(
	handler http.HandlerFunc,
	middleWares []MiddleWareFunc,
	w http.ResponseWriter,
	r *http.Request,
) {
	if len(middleWares) > 0 {
		var next func()
		next = func() {
			if executor.index >= len(middleWares) {
				handler(w, r)
			} else {
				middleWare := middleWares[executor.index]
				executor.index += 1
				middleWare(w, r, next)
			}
		}
		next()
	} else {
		handler(w, r)
	}
}

func (router *Router) requestExecutor(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		executor := middleWareExecutor{
			index: 0,
		}
		executor.execute(handler, router.middleWares, w, r)
	}
}

func (router *Router) HandleRoute(path string, handler http.HandlerFunc) {
	router.BaseServer.HandleFunc(path, router.requestExecutor(handler))
}

func NewRouter() Router {
	router := Router{
		BaseServer: http.NewServeMux(),
	}
	return router
}

func (router *Router) AddMiddleware(middleware MiddleWareFunc) {
	router.middleWares = append(router.middleWares, middleware)
}

func (router Router) ServeHttp(w http.ResponseWriter, r *http.Request) {

}
