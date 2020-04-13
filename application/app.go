package application

import (
	"fmt"
	"net/http"

	"github.com/abdullah-aghayan/urlShortener/logger"
	"github.com/go-chi/chi"
)

// Run start a web server on host with the specified port
func Run(host, port string) {
	addr := fmt.Sprintf("%s:%s", host, port)
	mux := route(chi.NewRouter())
	logger.Info(fmt.Sprintf("Server started on %s ...", addr))

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		logger.Error("Server failed to start. Please check logs for more details.", err)
	}
}
