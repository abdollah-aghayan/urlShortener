package url

import (
	"errors"
	"testing"

	restErr "github.com/abdullah-aghayan/urlShortener/utils/errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestValidateUrl(t *testing.T) {

	tt := []struct {
		name      string
		url       Url
		expextErr bool
		expect    *restErr.RestErr
	}{
		{
			name:      "empty url error",
			url:       Url{},
			expextErr: true,
			expect:    restErr.NewBadRequest("url is required"),
		},
		{
			name:      "url parse error",
			url:       Url{URL: "google.com"},
			expextErr: true,
			expect:    restErr.NewBadRequest("url is not valid!"),
		},
		{
			name:      "success",
			url:       Url{URL: "http://google.com"},
			expextErr: false,
			expect:    nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.url.ValidateUrl()

			if tc.expextErr {
				if err.Message != tc.expect.Message {
					t.Errorf("Expect %v got %v", tc.expect, err)
				}

				if err.Status() != tc.expect.Status() {
					t.Errorf("Expect %v got %v", tc.expect, err)
				}
				return
			}

			if err != nil {
				t.Errorf("Expected %v got %v", nil, err)
			}
		})
	}

}

func TestSaveUrl(t *testing.T) {
	// mock db
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("can not mock DB")
	}

	defer db.Close()
	dbx := sqlx.NewDb(db, "mockDB")
	storage := NewRepo(dbx)

	// Error on prepare
	url := New("111", "http://google.com")

	tt := []struct {
		name   string
		mock   func()
		err    string
		expect *Url
	}{
		{
			name: "preper error",
			mock: func() {
				mock.ExpectPrepare("INSERT INTO url").WillReturnError(errors.New("error"))
			},
			err:    "internal server error",
			expect: nil,
		},
		{
			name: "insert error",
			mock: func() {
				mock.ExpectPrepare("INSERT INTO url").ExpectExec().WillReturnError(errors.New("error"))
			},
			err:    "internal server error",
			expect: nil,
		},
		{
			name: "success",
			mock: func() {
				mock.ExpectPrepare("INSERT INTO url").ExpectExec().WillReturnResult(sqlmock.NewResult(1, 1))
			},
			err:    "",
			expect: &Url{"111", "http://google.com"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := storage.SaveURL(url)

			if err != nil {
				if err.Error() != tc.err {
					t.Fatalf("Expected %s error got %s", tc.err, err.Error())
				}
			}

			if tc.expect != nil {
				if res == nil {
					t.Fatalf("Expected %v got %v", tc.expect, res)
				}
			}
		})
	}

}

func TestGet(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("can not mock DB")
	}

	defer db.Close()
	dbx := sqlx.NewDb(db, "mockDB")
	storage := NewRepo(dbx)

	url := New("111", "http://google.com")

	tt := []struct {
		name string
		mock func()
		want *Url
		err  string
	}{
		{
			name: "prepar error",
			mock: func() {
				mock.ExpectPrepare("SELECT (.+) FROM urls").WillReturnError(errors.New("not found"))
			},
			want: nil,
			err:  "internal server error",
		},
		{
			name: "select error",
			mock: func() {
				mock.ExpectPrepare("SELECT (.+) FROM urls").ExpectQuery().WithArgs(url.ID).WillReturnError(errors.New("not found"))
			},
			want: nil,
			err:  "internal server error",
		},
		{
			name: "not found",
			want: nil,
			mock: func() {
				row := sqlmock.NewRows([]string{"id", "url"})
				mock.ExpectPrepare("SELECT (.+) FROM urls").ExpectQuery().WithArgs(url.ID).WillReturnRows(row)
			},
			err: "Can not find url",
		},
		{
			name: "success",
			want: &Url{"111", "http://google.com"},
			mock: func() {
				row := sqlmock.NewRows([]string{"id", "url"}).AddRow(url.ID, url.URL)
				mock.ExpectPrepare("SELECT (.+) FROM urls").ExpectQuery().WithArgs(url.ID).WillReturnRows(row)
			},
			err: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.mock()
			res, err := storage.GetURL(url.ID)
			if tc.want != nil {
				if tc.want.URL != res.URL {
					t.Fatalf("GetURL() expected %v got %v", tc.want, res)
				}
			}

			if err != nil {
				if tc.err != err.Error() {
					t.Fatalf("GetURL() expected %s error got %s", tc.err, err.Error())
				}
			}
		})
	}
}
