package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getInventoryHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func placeOrderHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func findOrderByIDHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func deleteOrderHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}
