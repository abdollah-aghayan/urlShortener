package repository

import (
	"fmt"
	"os"

	// just load sql dependeces
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Connect create a connection to database
func Connect() (*sqlx.DB, error) {
	dbConString := fmt.Sprintf("%s:%s@tcp(%s:%s)/urlShortener",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"))

	fmt.Println(dbConString)

	db, err := sqlx.Connect("mysql", dbConString)

	if err != nil {
		return nil, err
	}

	return db, err
}
