package routes

import "net/http"

type Router struct {
	Rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		Rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	routeExists, methodExists, handler := r.FindRoute(request.URL.Path, request.Method)

	if routeExists {
		if methodExists {
			handler(writer, request)
		} else {
			writer.WriteHeader(http.StatusMethodNotAllowed)
		}
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func (r *Router) FindRoute(path string, method string) (bool, bool, http.HandlerFunc) {
	_, routeExists := r.Rules[path]
	handler, methodExists := r.Rules[path][method]
	return routeExists, methodExists, handler
}
