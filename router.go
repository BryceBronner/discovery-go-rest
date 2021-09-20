package main

import "net/http"

type RequestHandlerFnc func(*RequestContext)

type RouteInfo struct {
	Path    string
	Method  string
	Handler RequestHandlerFnc
}

type RouterHandler struct {
	routes []RouteInfo
}

func (router *RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range router.routes {
		if route.Method != r.Method {
			continue
		}

		if route.Path == r.URL.Path {
			context := &RequestContext{w, r}
			route.Handler(context)
		}
	}
}

func (router *RouterHandler) GET(path string, handler RequestHandlerFnc) {
	router.routes = append(router.routes, RouteInfo{
		Path:    path,
		Method:  http.MethodGet,
		Handler: handler,
	})
}

func createRouter() *RouterHandler {
	return &RouterHandler{
		routes: []RouteInfo{},
	}
}
