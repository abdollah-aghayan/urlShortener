package main

import (
	"os"

	"github.com/abdullah-aghayan/urlShortener/application"
)

func main() {

	os.Setenv("ADDR", "localhost")
	os.Setenv("PORT", "8080")

	application.Run(os.Getenv("ADDR"), os.Getenv("PORT"))
}
