package application

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func TestRoute(t *testing.T) {
	srv := httptest.NewServer(route(chi.NewRouter()))
	defer srv.Close()

	res, err := http.Get(srv.URL)
	if err != nil {
		t.Fatal("Can not get home page")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("can not read response")
	}

	if string(data) == "" {
		t.Fatal("home page is empty")
	}
}
