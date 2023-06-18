package api

import (
	"net/http"

	"github.com/dannylee/url-ports-adapters/utils"
)

type msg map[string]any

func ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	res := msg{"error": message}

	err := utils.WriteJSON(w, status, res, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func ServerErrorResponse(w http.ResponseWriter, r *http.Request, message any) {
	ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func BadRequestResponse(w http.ResponseWriter, r *http.Request) {
	message := "Request is invalid"
	ErrorResponse(w, r, http.StatusBadRequest, message)
}
