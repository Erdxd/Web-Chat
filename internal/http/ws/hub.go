package http

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	Send   chan []byte
	RoomId int
}
type Hub struct {
	RoomId     map[int]map[*Client]bool
	Clients    map[*Client]bool
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
	Mu         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{RoomId: make(map[int]map[*Client]bool), Clients: make(map[*Client]bool), Broadcast: make(chan []byte), Register: make(chan *Client), Unregister: make(chan *Client)}
}
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Mu.Lock()
			h.Clients[client] = true
			h.Mu.Unlock()
		case client := <-h.Unregister:
			h.Mu.Lock()
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
			h.Mu.Unlock()
		case message := <-h.Broadcast:
			h.Mu.Lock()
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)

				}

			}
			h.Mu.Unlock()
		}
	}
}
func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()
	for message := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}
