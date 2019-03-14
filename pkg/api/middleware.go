package api

import (
	"net/http"

	"github.com/wexel-nath/pets/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

func middlewareWrapper(p path) httprouter.Handle {
	next := loggerWrapper(p.handler)
	return next
}

func loggerWrapper(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		logger.Info("%s %s", r.Method, r.URL.Path)
		next(w, r, p)
	}
}
