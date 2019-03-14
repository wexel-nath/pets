package api

import (
	"net/http"

	"github.com/wexel-nath/pets/pkg/logger"

	"github.com/julienschmidt/httprouter"
)

func healthzHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte(`{"status":"ok"}`))
	if err != nil {
		logger.Error(err)
	}
}
