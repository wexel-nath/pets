package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func addPetHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func updatePetHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func petGetHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	switch p.ByName("path") {
	case "findByStatus":
		findPetsByStatusHandler(w, r, p)
	default:
		findPetByIDHandler(w, r, p)
	}
}

func findPetsByStatusHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func findPetByIDHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func updatePetWithFormDataHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func deletePetHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func uploadPetImageHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}
