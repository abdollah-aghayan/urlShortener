package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {
	json, _ := json.Marshal(data)

	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
