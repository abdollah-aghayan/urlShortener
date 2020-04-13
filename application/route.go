package application

import (
	"github.com/abdullah-aghayan/urlShortener/handlers"
	"github.com/go-chi/chi"
)

// Route set urls
func route(mux *chi.Mux) *chi.Mux {

	mux.Get("/", handlers.Home)
	mux.Post("/url", handlers.CreateUrl)
	mux.Get("/{id}", handlers.GetUrl)

	return mux
}
