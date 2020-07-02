package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON ...
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if nil != err {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR ...
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if nil != err {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}