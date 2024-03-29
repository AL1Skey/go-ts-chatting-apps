package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	// encapsulate sql to db
	// lowercase mean private access on struct
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	// Open connection to postgres database on docker port
	// Get database and error message
	// syntax: postgres://<username>:<password>@<server-that-run>:<port-that-server-run>
	db, err := sql.Open("postgres", "postgres://root:password@localhost:5433")
	// return error of there is error message
	if err != nil {
		return nil, err
	}
	// return pointer to database struct that contain DB
	return &Database{db: db}, nil

}
