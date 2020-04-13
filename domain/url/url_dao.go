package url

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/abdullah-aghayan/urlShortener/logger"
	"github.com/jmoiron/sqlx"
)

const (
	querySaveUrl = "INSERT INTO urls (id, url) VALUES (?, ?)"
	queryGetUrl  = "SELECT * FROM urls WHERE id = ?"
)

var UrlStorage urlStorageInterface = &urlRepo{}

type urlStorageInterface interface {
	SaveURL(Url) (*Url, error)
	GetURL(string) (*Url, error)
}

type urlRepo struct {
	db *sqlx.DB
}

func NewRepo(db *sqlx.DB) urlStorageInterface {
	return &urlRepo{db}
}

// Initialize create a connection to database
func (storage *urlRepo) Initialize(user, password, host, port string) *sqlx.DB {
	dbConString := fmt.Sprintf("%s:%s@tcp(%s:%s)/urlShortener", user, password, host, port)

	var err error
	storage.db, err = sqlx.Connect("mysql", dbConString)

	if err != nil {
		fmt.Println("Can not connect to database ", err)
		os.Exit(2)
	}

	return storage.db
}

// SaveURL save client url to db
func (storage *urlRepo) SaveURL(url Url) (*Url, error) {

	stat, err := storage.db.Prepare(querySaveUrl)

	if err != nil {
		logger.Error("failed to prepar url", err)
		return nil, errors.New("internal server error")
	}

	defer stat.Close()

	_, err = stat.Exec(url.ID, url.URL)

	if err != nil {
		logger.Error("failed to save url", err)
		return nil, errors.New("internal server error")
	}

	return &url, nil
}

// GetURL fetch url form database by hash
func (storage *urlRepo) GetURL(id string) (*Url, error) {

	stat, err := storage.db.Preparex(queryGetUrl)

	if err != nil {
		logger.Error("failed to prepar url", err)
		return nil, errors.New("internal server error")
	}
	defer stat.Close()

	var url Url
	err = stat.Get(&url, id)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, errors.New("Can not find url")
		}
		logger.Error("database error", err)
		return nil, errors.New("internal server error")
	}

	return &url, nil
}
