package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wexel-nath/pets/pkg/logger"
	"github.com/wexel-nath/pets/pkg/pet"

	"github.com/julienschmidt/httprouter"
)

func addPetHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Error(err)
		response(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	var petModel pet.Pet
	err = json.Unmarshal(body, &petModel)
	if err != nil {
		logger.Error(err)
		response(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	newPet, err := pet.Insert(petModel)
	if err != nil {
		logger.Error(err)
		response(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response(w, http.StatusCreated, newPet, "")
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
