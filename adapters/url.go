package adapters

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/dannylee/url-ports-adapters/core/models"
)

type UrlAdapter struct {
	db  *sql.DB
	cfg Config
}

type Config struct {
	Port int
	Env  string
	DB   struct {
		Dsn            string
		MaxIdleTimeout string
	}
	Limiter bool
}

func New(cfg Config, db *sql.DB) *UrlAdapter {

	return &UrlAdapter{cfg: cfg, db: db}
}

func (u *UrlAdapter) QueryWithLong(url string) (*models.UrlModel, error) {
	query := `
		SELECT
			*
		FROM url
		WHERE
			long_url = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var data models.UrlModel

	err := u.db.QueryRowContext(ctx, query, url).Scan(
		&data.ID,
		&data.LongURL,
		&data.ShortURL,
	)
	if err != nil {
		log.Println("error querying with long url")
		return nil, err
	}

	return &data, nil
}

func (u *UrlAdapter) QueryWithShort(url string) (*models.UrlModel, error) {
	query := `
		SELECT
			*
		FROM url
		WHERE short_url = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var data models.UrlModel

	err := u.db.QueryRowContext(ctx, query, url).Scan(
		&data.ID,
		&data.ShortURL,
		&data.LongURL,
	)

	if err != nil {
		log.Println("error querying with short URL")
		return nil, err
	}

	return &data, nil
}

func (u *UrlAdapter) InsertURL(data models.UrlModel, shortUrl string) error {
	query := `
		INSERT INTO url (short_url, long_url)
		VALUES ($1, $2)
		RETURNING (id)
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := u.db.QueryRowContext(ctx, query, shortUrl, data.LongURL).Scan(&data.ID)
	if err != nil {
		return err
	}

	return nil
}
