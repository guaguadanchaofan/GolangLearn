package gee

import (
	"fmt"
	"net/http"
)

// 定义函数对象
type HandleFunc func(http.ResponseWriter, *http.Request)

// 传入的实例
type Engine struct {
	router map[string]HandleFunc
}

// 获取对象
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

// 接口函数
func (engine *Engine) addRoute(method string, pattern string, hander HandleFunc) {
	key := method + "-" + pattern
	engine.router[key] = hander
}

func (engine *Engine) GET(pattern string, hander HandleFunc) {
	engine.addRoute("GET", pattern, hander)
}

func (engine *Engine) POST(pattern string, hander HandleFunc) {
	engine.addRoute("POST", pattern, hander)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND : %s\n", r.URL)
	}
}
