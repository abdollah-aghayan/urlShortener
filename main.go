package main

import (
	"github.com/abdullah-aghayan/urlShortener/handlers"
	"github.com/abdullah-aghayan/urlShortener/repository"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {

	baseURL := os.Getenv("BASE_URL")
	if (baseURL == "") {
		fmt.Println("BASE_URL enviroment vaiable not available")
		os.Exit(1)
	}

	db, err := repository.Connect()

	if err != nil {
		fmt.Println("Can not connect to database ")
		os.Exit(2)
	}

	
	urlRepo := repository.URLRepo{Db: db, BaseURL: baseURL}

	urlHandler := handlers.NewURLHandler(urlRepo)

	handler := urlHandler.Route(chi.NewRouter())

	fmt.Println("Server started...")
	fmt.Println(http.ListenAndServe(":8080", handler))

}
