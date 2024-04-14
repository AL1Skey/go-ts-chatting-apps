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
	r.GET("/ws/get-room", websocketHandler.GetRoom)
	r.GET("/ws/join-room/:roomId", websocketHandler.JoinRoom)
	r.GET("/ws/get-client/:roomId", websocketHandler.GetClient)
}

func Start(addr string) error {
	return r.Run(addr)
}
