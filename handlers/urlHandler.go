package handlers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"

	"github.com/abdullah-aghayan/urlShortener/models"

	"github.com/abdullah-aghayan/urlShortener/repository"

	"github.com/go-chi/chi"
)

type URLHandler struct {
	repo repository.URLRepo
}

// NewURLHandler return new instance of UrlHandler
func NewURLHandler(repo repository.URLRepo) URLHandler {
	return URLHandler{repo}
}

// Route set urls
func (urlHandler URLHandler) Route(mux *chi.Mux) *chi.Mux {

	mux.Post("/url", urlHandler.create)
	mux.Get("/{hash}", urlHandler.get)

	// mux.Put("/url", urlHandler.update)
	// mux.Delete("/url", urlHandler.delete)
	return mux
}

func (urlHandler URLHandler) create(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var obj models.Url

	err := decoder.Decode(&obj)
	if err != nil {
		http.Error(w, "Invalid request body!", http.StatusBadRequest)
		return
	}

	// check url is required
	if len(obj.URL) == 0 {
		http.Error(w, "url is requiered!", http.StatusBadRequest)
		return
	}

	// url is standard lib
	_, err = url.ParseRequestURI(obj.URL)
	if err != nil {
		http.Error(w, "url is not currect!", http.StatusBadRequest)
		return
	}

	id := generateID()
	// check if id exist genarate new one

	// TODO check whether url exist o not

	dbErr := urlHandler.repo.SaveURL(obj, id)

	if dbErr != nil {
		fmt.Fprintf(w, "some thing went wrong with this data %s error is %s", obj, dbErr)
		return
	}

	success := map[string]string{"url": os.Getenv("BASE_URL") + ":8080/" + id}
	json, _ := json.Marshal(success)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func (urlHandler URLHandler) get(w http.ResponseWriter, r *http.Request) {

	hash := chi.URLParam(r, "hash")

	url, err := urlHandler.repo.GetURL(hash)

	if err != nil {
		http.Error(w, "404", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
	return
}

func (urlHandler URLHandler) update(w http.ResponseWriter, r *http.Request) {

}
func (urlHandler URLHandler) delete(w http.ResponseWriter, r *http.Request) {

}

// generateID genarate 6 charachter string for url
func generateID() string {
	var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, 6)
	for i := range s {
		s[i] = chars[rand.Intn(len(chars))]
	}

	return string(s)
}
