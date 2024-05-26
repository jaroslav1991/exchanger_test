package utils

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, code int, err error) {
	response := map[string]any{"error": fmt.Sprintf("%v", err.Error())}

	WriteSuccessResponse(w, code, response)
}

func WriteSuccessResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		response, err := json.Marshal(data)
		if err != nil {
			logrus.Error("cannot marshal response: ", err)
			return
		}
		_, err = w.Write(response)
		if err != nil {
			logrus.Error("cannot write response: ", err)
		}
	}
}
