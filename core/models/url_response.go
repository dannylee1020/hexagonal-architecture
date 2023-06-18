package models

import ()

type UrlModel struct {
	ID       string `json:"ID"`
	ShortURL string `json:"short_url"`
	LongURL  string `json:"long_url"`
}
