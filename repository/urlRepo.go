package repository

import (
	"errors"

	"github.com/abdullah-aghayan/urlShortener/models"

	"github.com/jmoiron/sqlx"
)

// URLRepo contain repository for url
type URLRepo struct {
	BaseURL string
	Db      *sqlx.DB
}

// SaveURL save client url to db
func (urlRepo URLRepo) SaveURL(url models.Url, id string) error {

	_, err := urlRepo.Db.Exec("INSERT INTO urls (id, url) VALUES (?, ?)", id, url.URL)

	if err != nil {
		return err
	}

	return nil
}

// GetURL fetch url form database by hash
func (urlRepo URLRepo) GetURL(hash string) (string, error) {

	url := models.Url{}
	urlRepo.Db.Get(&url, "Select * from urls where id = ?", hash)

	if url == (models.Url{}) {
		return "", errors.New("can not find url")
	}

	return url.URL, nil
}
