package users

import (
	"context"
	"database/sql"
)

// DBTX is an interface that defines a set of methods for executing SQL queries and transactions.
// It can be used to interact with a database in a type-safe way.
type DBTX interface {
	// Exec executes a query without returning any rows.
	// It is typically used for insert, update, or delete operations.
	Exec(query string, args ...interface{}) (sql.Result, error)
	// Query executes a query that returns rows, and returns a *sql.Rows result.
	// It can be used to retrieve data from the database.
	Query(query string, args ...interface{}) (*sql.Rows, error)
	// QueryRow executes a query that returns at most one row.
	// It can be used to retrieve a single row of data from the database.
	QueryRow(query string, args ...interface{}) *sql.Row
	// QueryContext is similar to Query, but takes a context.Context as its first argument.
	// It is used to cancel the query if the context is canceled.
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	// QueryRowContext is similar to QueryRow, but takes a context.Context as its first argument.
	// It is used to cancel the query if the context is canceled.
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	// Prepare creates a prepared statement for the given query.
	// Prepared statements can be used to improve the performance of queries that are executed multiple times.
	Prepare(query string) (*sql.Stmt, error)
	// PrepareContext is similar to Prepare, but takes a context.Context as its first argument.
	// It is used to cancel the preparation if the context is canceled.
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	// Begin starts a new transaction.
	Begin() (*sql.Tx, error)
	// BeginTx is similar to Begin, but takes a context.Context and a sql.TxOptions as its arguments.
	// It is used to start a transaction with specific options.
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	// Commit commits the current transaction.
	Commit() error
	// Rollback rolls back the current transaction.
	Rollback() error
	// Close closes the database connection.
	Close() error
}

// repository is a struct that contains a DBTX field, which is used to interact with the database.
type repository struct {
	db DBTX
}

// NewRepository is a function that takes a DBTX as its argument and returns a *repository.
// It is used to create a new instance of the repository.
func NewRepository(db DBTX) *repository {
	return &repository{db: db}
}

// CreateUser is a method that takes a context.Context and a pointer to a User struct as its arguments.
// It is used to insert a new user into the database.
// It returns a pointer to the User struct and an error.
// syntax: func (<receiver AKA inheritence>) <name function>(<parameter>) (<return value>,<return value>)
func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	// lastInsertId is used to store the ID of the inserted user.
	var lastInsertId int64
	// query is a string that contains the SQL query for inserting a new user.
	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) returning id"
	// err is used to store the error returned by the QueryRowContext method.
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertId)

	if err != nil {
		// If there is an error, it is returned to the caller.
		return nil, err
	}

	// The ID of the inserted user is set to the User struct.
	user.ID = int64(lastInsertId)
	// The User struct and nil error are returned to the caller.
	return user, nil
}
