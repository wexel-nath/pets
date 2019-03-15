package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wexel-nath/pets/pkg/logger"
	"github.com/wexel-nath/pets/pkg/user"

	"github.com/julienschmidt/httprouter"
)

type UserApiModel struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Status    int64  `json:"userStatus"`
}

func NewUserApiModel(u user.User) UserApiModel {
	return UserApiModel{
		ID:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
		Phone:     u.Phone,
		Status:    u.Status,
	}
}

func (model UserApiModel) ToUser() user.User {
	return user.User{
		ID:        model.ID,
		Username:  model.Username,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.Email,
		Password:  model.Password,
		Phone:     model.Phone,
		Status:    model.Status,
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		logger.Error(err)
		response(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	var model UserApiModel
	err = json.Unmarshal(body, &model)
	if err != nil {
		logger.Error(err)
		response(w, http.StatusBadRequest, nil, err.Error())
		return
	}

	u, err := user.Insert(model.ToUser())
	if err != nil {
		logger.Error(err)
		response(w, http.StatusInternalServerError, nil, err.Error())
		return
	}

	response(w, http.StatusCreated, NewUserApiModel(u), "")
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
