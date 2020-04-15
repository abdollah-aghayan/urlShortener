package main

import (
	"os"

	"github.com/abdullah-aghayan/urlShortener/application"
)

func main() {

	application.Run(os.Getenv("ADDR"), os.Getenv("PORT"))
}
