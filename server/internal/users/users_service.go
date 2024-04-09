// Package users provides a user service that handles user creation, retrieval, and deletion.
package users

import (
	"context"              // Provides a context object that carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.
	"server/internal/util" // Provides utility functions for the application.
	"strconv"              // Provides functions for converting between string and numeric types.
	"time"                 // Provides functionality for measuring and displaying time.
)

// service is a struct that contains a Repository and a timeout duration.
type service struct {
	Repository               // Represents a database or other storage system that the user service can use to store and retrieve user data.
	timeout    time.Duration // Represents the maximum amount of time that the user service will wait for a database operation to complete.
}

// NewService creates a new user service with the given repository and timeout duration.
func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second, // Sets the timeout duration to 2 seconds.
	}
}

// CreateUser creates a new user in the system.
func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout) // Creates a new context that is a copy of the parent context but has a timeout set to the value of s.timeout.
	defer cancel()                                   // Cancels the context when the function returns.

	// Hashes the password provided in the request object using the util.HashPassword function.
	hashpw, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Creates a new User object with the username, hashed password, and email from the request object.
	u := &User{
		Username: req.Username,
		Password: hashpw,
		Email:    req.Email,
	}

	// Saves the new user to the database using the Repository object.
	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	// Creates a new response object with the ID, username, and email of the newly created user.
	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}
	return res, nil
}
