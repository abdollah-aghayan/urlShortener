package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/abdullah-aghayan/urlShortener/domain/url"
	"github.com/abdullah-aghayan/urlShortener/logger"
	"github.com/abdullah-aghayan/urlShortener/service"
	"github.com/abdullah-aghayan/urlShortener/utils/errors"
	"github.com/abdullah-aghayan/urlShortener/utils/response"

	"github.com/go-chi/chi"
)

func CreateUrl(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var url url.Url
	err := decoder.Decode(&url)
	if err != nil {
		resErr := errors.NewBadRequest("Can not read request body")
		response.JSON(w, resErr.Status(), resErr)
		return
	}

	res, restErr := service.UrlService.SaveUrl(url.URL)
	if restErr != nil {
		response.JSON(w, restErr.Status(), restErr)
		return
	}

	data := map[string]string{"id": res.ID}

	response.JSON(w, http.StatusCreated, data)
	return
}

func GetUrl(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		err := errors.NewBadRequest("id is required")
		response.JSON(w, err.Status(), err)
		return
	}

	url, err := service.UrlService.GetUrl(id)
	if err != nil {
		logger.Debug(err.Error())
		response.JSON(w, err.Status(), err)
		return
	}

	http.Redirect(w, r, url.URL, http.StatusMovedPermanently)
	return
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h4>welcome to Url Shortener</h4>")
	return
}
