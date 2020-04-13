package url

import (
	"net/url"

	"github.com/abdullah-aghayan/urlShortener/utils/errors"
	_ "github.com/go-sql-driver/mysql"
)

type Url struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

func (u Url) ValidateUrl() *errors.RestErr {
	// check url is required
	if len(u.URL) == 0 {
		return errors.NewBadRequest("url is required")
	}

	// url is standard lib
	_, err := url.ParseRequestURI(u.URL)
	if err != nil {
		return errors.NewBadRequest("url is not valid!")
	}

	return nil
}

func New(id string, url string) Url {
	return Url{
		id,
		url,
	}
}
