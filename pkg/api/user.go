package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wexel-nath/pets/pkg/logger"
	"github.com/wexel-nath/pets/pkg/user"

	"github.com/julienschmidt/httprouter"
)

func createUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Error(err)
		response(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	var userModel user.User
	err = json.Unmarshal(body, &userModel)
	if err != nil {
		logger.Error(err)
		response(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	newUser, err := user.Insert(userModel)
	if err != nil {
		logger.Error(err)
		response(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response(w, http.StatusCreated, newUser, "")
}

func createUserWithArrayHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func createUserWithListHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func userGetHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	switch p.ByName("path") {
	case "login":
		userLoginHandler(w, r, p)
	case "logout":
		userLogoutHandler(w, r, p)
	default:
		getUserByUsernameHandler(w, r, p)
	}
}

func userLoginHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func userLogoutHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func getUserByUsernameHandler(w http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	username := p.ByName("path")
	logger.Info(username)
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func updateUserHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}

func deleteUserHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	// TODO
	w.WriteHeader(http.StatusNotImplemented)
}
