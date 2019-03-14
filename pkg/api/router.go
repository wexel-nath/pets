package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetRouter() *httprouter.Router {
	router := httprouter.New()

	for _, path := range getPaths() {
		router.Handle(path.method, path.pattern, middlewareWrapper(path))
	}

	return router
}

type path struct {
	method  string
	pattern string
	handler httprouter.Handle
}

func getPaths() []path {
	return []path{
		{
			method:  http.MethodGet,
			pattern: "/healthz",
			handler: healthzHandler,
		},
	}
}
