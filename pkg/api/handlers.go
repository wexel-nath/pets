package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/wexel-nath/pets/pkg/logger"
	"github.com/wexel-nath/pets/pkg/user"

	"github.com/julienschmidt/httprouter"
)

func healthzHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte(`{"status":"ok"}`))
	if err != nil {
		logger.Error(err)
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
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

type Response struct {
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
}

func response(w http.ResponseWriter, status int, result interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := Response{
		Result:  result,
		Message: message,
	}

	mJSON, err := json.Marshal(resp)
	logger.LogIfErr(err)

	_, err = w.Write(mJSON)
	logger.LogIfErr(err)
}
