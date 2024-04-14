package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// writeMessage is a method of the Client struct that writes messages to the client's WebSocket connection.
// It reads messages from the client's Message channel and sends them to the client's WebSocket connection.
// It closes the WebSocket connection when the Message channel is closed.
func (c *Client) writeMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		msg, ok := <-c.Message
		if !ok {
			return
		}
		c.Conn.WriteJSON(msg)
	}
}

// readMessage is a method of the Client struct that reads messages from the client's WebSocket connection.
// It sends the received messages through the Hub's Broadcast channel.
// It unregisters the client from the Hub and closes the WebSocket connection when an error occurs.
func (c *Client) readMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error %v", err)
			}
		}
		m := &Message{
			Content:  string(msg),
			RoomID:   c.RoomId,
			Username: c.Username,
		}
		hub.Broadcast <- m
	}
}

func (hub *Handler) GetClient(c *gin.Context) {
	var client []ClientResponse

	roomId := c.Param("roomId")

	if _, ok := hub.hub.Rooms[roomId]; ok {
		client = make([]ClientResponse, 0)
		c.JSON(http.StatusOK, client)
	}

	for _, c := range hub.hub.Rooms[roomId].Clients {
		client = append(client, ClientResponse{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, client)
}
