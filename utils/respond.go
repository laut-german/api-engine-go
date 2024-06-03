package utils

import (
	"encoding/json"
	"net/http"
	"web-engine-go/middleware"
)

func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) error {
	if e, ok := data.(error); ok {
		var temp = new(struct {
			Status string `json:"status"`
			Error  string `json:"error"`
		})

		temp.Status = "error"
		temp.Error = e.Error()
		data = temp
	}
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	middleware.LogRequest(r, status)
	return nil
}