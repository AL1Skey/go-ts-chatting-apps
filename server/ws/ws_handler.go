package ws

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Handler struct contains a reference to the Hub, which manages the rooms and clients.
type Handler struct {
	hub *Hub
}

// NewHandler is a constructor function that creates a new Handler instance with the given Hub.
func NewHandler(hub *Hub) *Handler {
	return &Handler{hub: hub}
}

// CreateRoom is a Gin HTTP handler function that creates a new room with the given ID and name.
// It adds the new room to the Hub's Rooms map.
func (hub *Handler) CreateRoom(c *gin.Context) {
	var request CreateRoomReq
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, request)

	hub.hub.Rooms[request.ID] = &Room{
		ID:      request.ID,
		Name:    request.Name,
		Clients: make(map[string]*Client),
	}
}

// upgrader is a Gorilla WebSocket Upgrader instance that is used to upgrade HTTP connections to WebSocket connections.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// JoinRoom is a Gin HTTP handler function that upgrades the HTTP connection to a WebSocket connection and adds the new client to the specified room.
// It also broadcasts a message to the room indicating that a new user has joined.
func (hub *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Route /ws/JoinRoom/:roomId?userId=1&username=user
	roomID := c.Param("roomId")
	clientID := c.Query("userId")
	username := c.Query("username")

	client := &Client{
		ID:       clientID,
		Username: username,
		RoomId:   roomID,
		Conn:     conn,
		Message:  make(chan *Message, 10), // Buffer Message of 10
	}

	message := &Message{
		Username: username,
		RoomID:   roomID,
		Content:  fmt.Sprintf("New user are joining the room %s", roomID),
	}

	// Register a new Client through the register channel
	hub.hub.Register <- client
	// Broadcast the message
	hub.hub.Broadcast <- message

	go client.writeMessage()
	client.readMessage(hub.hub)
}
