package ports

import (
	"github.com/dannylee/url-ports-adapters/core/models"
)

type UrlService interface {
	QueryWithLong(url string) (*models.UrlModel, error)
	QueryWithShort(url string) (*models.UrlModel, error)
	InsertURL(urlData models.UrlModel, shortUrl string) error
}
