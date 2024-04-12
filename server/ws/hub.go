package ws

import "fmt"

// NewHub is a constructor function that creates a new Hub instance with empty Rooms, Register, Unregister, and Broadcast channels.
func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

// Run is a method of the Hub struct that runs the Hub's main loop.
// It listens for new clients on the Register channel, unregisters clients on the Unregister channel, and broadcasts messages on the Broadcast channel.
func (h *Hub) Run() {
	for {
		select {
		// Register is a channel that receives new clients.
		// If the client's room exists, it adds the client to the room's Clients map.
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomId]; ok {
				r := h.Rooms[cl.RoomId]
				if _, ok := h.Rooms[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
				}
			}
		// Unregister is a channel that receives clients to be unregistered.
		// If the client's room exists, it removes the client from the room's Clients map.
		// If the room is empty after unregistering the client, it broadcasts a message indicating that the user left the chat.
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomId]; ok {
				if _, ok := h.Rooms[cl.ID]; ok {
					if len(h.Rooms[cl.RoomId].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  fmt.Sprintf("user %s left the chat", cl.ID),
							RoomID:   cl.RoomId,
							Username: cl.Username,
						}
					}

					delete(h.Rooms[cl.RoomId].Clients, cl.ID)
				}
			}
		// Broadcast is a channel that receives messages to be broadcasted.
		// If the message's room exists, it sends the message to all clients in the room.
		case msg := <-h.Broadcast:
			if _, ok := h.Rooms[msg.RoomID]; ok {
				for _, cl := range h.Rooms[msg.RoomID].Clients {
					cl.Message <- msg
				}
			}
		}
	}
}
