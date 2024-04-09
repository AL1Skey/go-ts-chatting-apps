package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Database struct encapsulates an SQL database connection.
type Database struct {
	// db is the underlying SQL.DB instance.
	db *sql.DB
}

// NewDatabase opens a connection to a Postgres database and returns a pointer to a new Database instance.
// The connection string should be in the format "postgres://<username>:<password>@<server-that-run>:<port-that-server-run>".
// If an error occurs during the connection process, it is returned as the second value.
// Otherwise, a pointer to a new Database instance containing the connection is returned.
func NewDatabase() (*Database, error) {
	// Open a connection to the Postgres database on the local machine at port 5433.
	// The username and password are assumed to be "root" and "password" respectively.
	// If an error occurs during the connection process, it is returned as the second value.
	// Otherwise, a pointer to a new Database instance containing the connection is returned.
	db, err := sql.Open("postgres", "postgres://root:password@localhost:5433/go-chat?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &Database{db: db}, nil
}

// Close closes the underlying SQL.DB instance.
func Close(d *Database) error {
	return d.db.Close()
}

// GetDB returns the underlying SQL.DB instance.
func (d *Database) GetDB() *sql.DB {
	return d.db
}
