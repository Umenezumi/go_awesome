package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	*router
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContext(writer, request)
	e.handle(c)
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}

func (e *Engine) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, e))
}
