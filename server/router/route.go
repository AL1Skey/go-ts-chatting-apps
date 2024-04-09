package router

import (
	"server/internal/users"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// NewRouter creates a new gin router.
func InitHandler(userHandler *users.Handler) {
	r = gin.Default()

	r.POST("/register", userHandler.CreateUser)

}

func Start(addr string) error {
	return r.Run(addr)
}
