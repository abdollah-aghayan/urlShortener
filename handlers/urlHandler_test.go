package handlers

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/abdullah-aghayan/urlShortener/domain/url"
	"github.com/abdullah-aghayan/urlShortener/service"
	"github.com/abdullah-aghayan/urlShortener/utils/errors"
	"github.com/go-chi/chi"
)

var (
	getUrl  func(string) (*url.Url, *errors.RestErr)
	saveUrl func(string) (*url.Url, *errors.RestErr)
)

type mockService struct{}

func (s *mockService) GetUrl(id string) (*url.Url, *errors.RestErr) {
	return getUrl(id)
}

func (s *mockService) SaveUrl(url string) (*url.Url, *errors.RestErr) {
	return saveUrl(url)
}

func TestCreateHandler(t *testing.T) {
	id := "hrd32"
	tt := []struct {
		name   string
		mock   func(u string) (*url.Url, *errors.RestErr)
		input  string
		out    string
		status int
	}{
		{
			name: "unprocessable body error",
			mock: func(u string) (*url.Url, *errors.RestErr) {
				return nil, errors.NewBadRequest("Can not read request body")
			},
			input:  `{"":"google2.com"}`,
			out:    `{"message":"Can not read request body"}`,
			status: http.StatusBadRequest,
		},
		{
			name: "domain err",
			mock: func(u string) (*url.Url, *errors.RestErr) {
				return nil, errors.NewInternalError("error from domain")
			},
			input:  `{"url":"google2.com"}`,
			out:    `{"message":"error from domain"}`,
			status: http.StatusInternalServerError,
		},
		{
			name:  "success",
			input: `{"url":"http://google2.com"}`,
			mock: func(u string) (*url.Url, *errors.RestErr) {
				return &url.Url{id, ""}, nil
			},
			out:    `{"id":"` + id + `"}`,
			status: http.StatusCreated,
		},
		{
			name:  "bad request body",
			input: `rl":"http://google2.com"}`, // wrong input
			mock: func(u string) (*url.Url, *errors.RestErr) { // We don't need mock here
				return nil, nil
			},
			out:    `{"message":"Can not read request body"}`,
			status: http.StatusBadRequest,
		},
	}

	service.UrlService = &mockService{}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			// just need value
			saveUrl = tc.mock

			req, err := http.NewRequest(http.MethodPost, "/url", bytes.NewBufferString(tc.input))

			if err != nil {
				t.Error("Error on create request")
			}

			rec := httptest.NewRecorder()

			CreateUrl(rec, req)

			res := rec.Result()

			defer res.Body.Close()

			resData, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal("Got error during parsing body")
			}

			if string(resData) != tc.out {
				t.Fatalf("Expected %s got %s", tc.out, string(resData))
			}

			if tc.status != res.StatusCode {
				t.Fatalf("Expected %d got %d", tc.status, res.StatusCode)
			}
		})
	}
}

func TestGetHandler(t *testing.T) {
	tt := []struct {
		name   string
		mock   func(u string) (*url.Url, *errors.RestErr)
		id     string
		expect string
		status int
	}{
		{
			name: "id error",
			mock: func(u string) (*url.Url, *errors.RestErr) {
				return nil, errors.NewInternalError("url not found error from service")
			},
			id:     "",
			expect: `{"message":"id is required"}`,
			status: http.StatusBadRequest,
		},
		{
			name: "url not found",
			mock: func(u string) (*url.Url, *errors.RestErr) {
				return nil, errors.NewInternalError("url not found error from service")
			},
			id:     "1234",
			expect: `{"message":"url not found error from service"}`,
			status: http.StatusInternalServerError,
		},
		{
			name: "success",
			id:   "1f2h3",
			mock: func(u string) (*url.Url, *errors.RestErr) {
				return &url.Url{"1f2h3", "http://google.com"}, nil
			},
			expect: "http://google.com",
			status: http.StatusMovedPermanently,
		},
	}

	service.UrlService = &mockService{}

	for _, tc := range tt {

		t.Run(tc.name, func(t *testing.T) {
			// just need value
			getUrl = tc.mock

			req, err := http.NewRequest(http.MethodGet, "/url/"+tc.id, nil)

			if err != nil {
				t.Error("Error on create request")
			}

			rec := httptest.NewRecorder()

			// Resolve request context
			rctx := chi.NewRouteContext()
			rctx.URLParams.Add("id", tc.id)
			req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

			GetUrl(rec, req)

			res := rec.Result()

			defer res.Body.Close()

			if tc.status == http.StatusMovedPermanently {
				l, _ := res.Location()
				if l.String() != tc.expect {
					t.Fatalf("Expected %s got %s", tc.expect, l.String())
				}
			} else {
				resData, err := ioutil.ReadAll(res.Body)
				if err != nil {
					t.Fatal("Got error during parsing body")
				}

				if string(resData) != tc.expect {
					t.Fatalf("Expected %s got %s", tc.expect, string(resData))
				}
			}

			if tc.status != res.StatusCode {
				t.Fatalf("Expected %d got %d", tc.status, res.StatusCode)
			}
		})
	}
}

func TestHomePage(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)

	if err != nil {
		t.Error("Error on create request")
	}

	rec := httptest.NewRecorder()
	Home(rec, req)

	res := rec.Result()

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Got error during parsing body")
	}

	if !strings.Contains(string(body), "welcome to Url Shortener") {
		t.Fatal("Home page error")
	}
}
