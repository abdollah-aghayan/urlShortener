package service

import (
	"math/rand"
	"time"

	"github.com/abdullah-aghayan/urlShortener/domain/url"
	"github.com/abdullah-aghayan/urlShortener/utils/errors"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var UrlService urlServiceInterface = &urlService{}

type urlServiceInterface interface {
	GetUrl(string) (*url.Url, *errors.RestErr)
	SaveUrl(string) (*url.Url, *errors.RestErr)
}

type urlService struct{}

func (s *urlService) GetUrl(id string) (*url.Url, *errors.RestErr) {

	url, err := url.UrlStorage.GetURL(id)
	if err != nil {
		return nil, errors.NewNotFoundRequest("url not found")
	}

	return url, nil
}

func (s *urlService) SaveUrl(in string) (*url.Url, *errors.RestErr) {
	var u url.Url
	u.URL = in

	err := u.ValidateUrl()
	if err != nil {
		return nil, err
	}

	u.ID = generateID()
	// check if id exist genarate new one

	// TODO check whether url exist o not
	res, dbErr := url.UrlStorage.SaveURL(u)

	if dbErr != nil {
		// TODO log error
		return nil, errors.NewInternalError("Something went wrong please try later")
	}

	return res, nil
}

// generateID genarate 6 character string for url
func generateID() string {
	var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())

	s := make([]rune, 6)
	for i := range s {
		s[i] = chars[rand.Intn(len(chars))]
	}

	return string(s)
}
