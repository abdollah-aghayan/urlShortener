package service

import (
	"errors"
	"net/http"
	"testing"

	"github.com/abdullah-aghayan/urlShortener/domain/url"
)

var (
	saveUrl func(u url.Url) (*url.Url, error)
	getUrl  func(string) (*url.Url, error)
)

type dbMock struct{}

func (m *dbMock) SaveURL(u url.Url) (*url.Url, error) {
	return saveUrl(u)
}

func (m *dbMock) GetURL(id string) (*url.Url, error) {
	return getUrl(id)
}

func TestGetUrlSuccess(t *testing.T) {
	url.UrlStorage = &dbMock{}

	getUrl = func(id string) (*url.Url, error) {
		return &url.Url{
			id,
			"http://google.com",
		}, nil
	}

	u, err := UrlService.GetUrl("123")

	if err != nil {
		t.Errorf("Expected nil got %v", err)
	}

	if u == nil {
		t.Errorf("Expected %v got nil", u)
	}

	if u.ID != "123" {
		t.Errorf("Expected %s got %s", "123", u.ID)
	}
}

func TestGetUrlError(t *testing.T) {
	url.UrlStorage = &dbMock{}

	/* mock db */
	getUrl = func(id string) (*url.Url, error) {
		return nil, errors.New("not found")
	}

	u, err := UrlService.GetUrl("123")

	if err.Status() != http.StatusNotFound {
		t.Errorf("Expected %d got %d", http.StatusNotFound, err.Status())
	}

	if u != nil {
		t.Errorf("Expected %v got nil", u)
	}

}

func TestSaveUrlSuccess(t *testing.T) {
	url.UrlStorage = &dbMock{}

	saveUrl = func(u url.Url) (*url.Url, error) {
		return &u, nil
	}

	u, err := UrlService.SaveUrl("http://google.com")

	if err != nil {
		t.Errorf("Expected nil got %v", err)
	}

	if u.ID == "" {
		t.Error("url id can not be empty")
	}

}

func TestSaveUrlError(t *testing.T) {
	url.UrlStorage = &dbMock{}

	saveUrl = func(u url.Url) (*url.Url, error) {
		return nil, errors.New("error on save")
	}

	_, err := UrlService.SaveUrl("http://google.com")

	if err == nil {
		t.Errorf("Expected %v got nil", err)
	}

	if err.Status() != http.StatusInternalServerError {
		t.Errorf("Expected %d status code got %d", http.StatusInternalServerError, err.Status())
	}
}

func TestSaveUrlValidationError(t *testing.T) {
	url.UrlStorage = &dbMock{}

	saveUrl = func(u url.Url) (*url.Url, error) {
		return nil, errors.New("error on save")
	}

	_, err := UrlService.SaveUrl("")

	if err == nil {
		t.Errorf("Expected %v got nil", err)
	}

	if err.Status() != http.StatusBadRequest {
		t.Errorf("Expected %d status code got %d", http.StatusBadRequest, err.Status())
	}
}
