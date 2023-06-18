package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Router(u *UrlController) http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", u.HealthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/shorten", u.ShortenUrlHandler)
	router.Handle(http.MethodGet, "/v1/redirect/:url", u.RedirectHandler)

	return router
}
