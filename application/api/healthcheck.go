package api

import (
	"encoding/json"
	"net/http"
)

func (u *UrlController) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status": "OK",
	}

	res, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
