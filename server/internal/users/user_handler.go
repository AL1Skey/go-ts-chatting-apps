// package users
//
// This package contains the logic for handling user-related operations.
//
// Important:
//  - "net/http" package is used for handling HTTP requests and responses.
//  - "github.com/gin-gonic/gin" package is used for creating RESTful APIs.
//
// Handler struct:
//  - It is a struct that holds a reference to the Service interface.
//
// NewHandler function:
//  - It creates a new instance of the Handler struct and returns a pointer to it.
//
// CreateUser method:
//  - It handles the creation of a new user.
//  - It takes a pointer to a gin.Context object as a parameter, which represents the HTTP request.
//  - It binds the JSON data from the request body to a CreateUserReq struct.
//  - If there is an error while binding the JSON data, it returns a HTTP response with status code 400 (Bad Request) and an error message.
//  - It calls the CreateUser method of the Service interface with the context and the CreateUserReq struct as parameters.
//  - If there is an error while calling the CreateUser method, it returns a HTTP response with status code 500 (Internal Server Error) and an error message.
//  - If the CreateUser method returns a result successfully, it returns a HTTP response with status code 200 (OK) and the result.

package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler struct
type Handler struct {
	Service
}

// NewHandler function
func NewHandler(s Service) *Handler {
	return &Handler{Service: s}
}

// CreateUser method
func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}
