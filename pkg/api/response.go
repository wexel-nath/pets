package api

import (
	"encoding/json"
	"net/http"

	"github.com/wexel-nath/pets/pkg/logger"
)

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
