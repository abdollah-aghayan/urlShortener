package response_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abdullah-aghayan/urlShortener/utils/response"
)

func TestJSON(t *testing.T) {

	rec := httptest.NewRecorder()

	input := map[string]string{"test": "test"}
	response.JSON(rec, http.StatusOK, input)

	res, err := ioutil.ReadAll(rec.Result().Body)
	if err != nil {
		t.Fatal("can not read from response")
	}

	if rec.Result().StatusCode != http.StatusOK {
		t.Fatalf("expected status code %d got %d", http.StatusOK, rec.Result().StatusCode)
	}
	expected := `{"test":"test"}`
	if string(res) != expected {
		t.Fatalf("expected %s got %s", expected, res)
	}
}
