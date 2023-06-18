package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/dannylee/url-ports-adapters/core/models"
	"github.com/dannylee/url-ports-adapters/core/ports"
	"github.com/dannylee/url-ports-adapters/utils"
)

type UrlController struct {
	urlService ports.UrlService
}

func NewUrlController(urlService ports.UrlService) *UrlController {
	return &UrlController{urlService: urlService}
}

func (u *UrlController) ShortenUrlHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ShortURL string `json:"short_url"`
		LongURL  string `json:"long_url"`
	}

	err := utils.ReadJSON(w, r, &input)
	if err != nil {
		BadRequestResponse(w, r)
		return
	}

	data := models.UrlModel{
		LongURL: input.LongURL,
	}

	_, err = u.urlService.QueryWithLong(input.LongURL)
	if err == nil {
		ServerErrorResponse(w, r, "Shortened url already exists in the DB")
		return
	}

	hash := utils.GenerateShortURL()

	err = u.urlService.InsertURL(data, hash)
	if err != nil {
		ServerErrorResponse(w, r, "Error inserting data into DB")
		return
	}

	data.ShortURL = hash
	err = utils.WriteJSON(w, 200, data, nil)
	if err != nil {
		ServerErrorResponse(w, r, "Error writing response")
		return
	}
}

func (u *UrlController) RedirectHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	shortUrl := params.ByName("url")

	urlData, err := u.urlService.QueryWithShort(shortUrl)
	if err != nil {
		ServerErrorResponse(w, r, "Error querying data from DB")
	}

	err = utils.WriteJSON(w, 200, urlData, nil)
	if err != nil {
		ServerErrorResponse(w, r, "Error writing response")
		return
	}

	utils.OpenWebPage(urlData.LongURL)
}
