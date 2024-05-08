package utils

import (
	"encoding/json"
	"net/http"
)

func SendJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")

	str, err := json.Marshal(data)
	if err != nil {
		message := "Error converting users to JSON"
		SendError(w, status, message, data)
		return
	}
	w.WriteHeader(status)
	w.Write(str)
}
