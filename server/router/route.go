package router

import (
	"server/internal/users"
	"server/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

// NewRouter creates a new gin router.
func InitHandler(userHandler *users.Handler, websocketHandler *ws.Handler) {
	r = gin.Default()

	// Users Routings
	r.POST("/register", userHandler.CreateUser)
	r.POST("/login", userHandler.LoginUser)
	r.GET("/logout", userHandler.LogoutUser)

	// Rooms Routings
	r.POST("/ws/create-room", websocketHandler.CreateRoom)
	r.GET("/ws/join-room", websocketHandler.JoinRoom)
}

func Start(addr string) error {
	return r.Run(addr)
}
