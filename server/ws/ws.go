package ws

import "github.com/gorilla/websocket"

// Room Section
type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[string]*Room `json:"rooms"`
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

// Peer2Peer Section
type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomId   string `json:"room_id"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	Username string `json:"username"`
	RoomID   string `json:"room_id"`
}
