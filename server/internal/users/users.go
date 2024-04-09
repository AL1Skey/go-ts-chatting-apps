// The `users` package contains code for working with a database.
package users

import "context"

// A User represents a single user in the database, with a unique ID, username, email, and password.
type User struct {
	// ID is a unique number for the user.
	ID int64 `json:"id" db:"id"`

	// Username is the user's chosen name to login.
	Username string `json:"username" db:"username"`

	// Email is the user's email address.
	Email string `json:"email" db:"email"`

	// Password is the user's chosen password.
	Password string `json:"password" db:"password"`
}

// Repository is an interface that represents a thing that can do different things to the `users` table.
type Repository interface {
	// CreateUser takes a new user and a special tag that says what computer is doing the command.
	// It then sends a command to the database to add the new user.
	CreateUser(contextTag context.Context, newUser *User) (*User, error)
}

// Service is an interface that represents a thing that can do different things to the `users` table.
type Service interface {
	// CreateUser takes a new user and a special tag that says what computer is doing the command.
	// It then sends a command to the database to add the new user.
	CreateUser(ctx context.Context, user *CreateUserReq) (*CreateUserRes, error)
}

// CreateUserReq is a struct that represents a request to create a new user.
type CreateUserReq struct {
	// Username is the user's chosen name to login.
	Username string `json:"username" db:"username"`

	// Password is the user's chosen password.
	Password string `json:"password" db:"password"`

	// Email is the user's email address.
	Email string `json:"email" db:"email"`
}

// CreateUserRes is a struct that represents a response to a request to create a new user.
type CreateUserRes struct {
	// ID is a unique number for the user.
	ID string `json:"id" db:"id"`

	// Username is the user's chosen name to login.
	Username string `json:"username" db:"username"`

	// Email is the user's email address.
	Email string `json:"email" db:"email"`
}
