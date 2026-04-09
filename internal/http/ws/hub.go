package http

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	Data   []byte
	RoomId int
}
type Client struct {
	Conn   *websocket.Conn
	Send   chan []byte
	RoomId int
}
type Hub struct {
	RoomId     map[int]map[*Client]bool
	Clients    map[*Client]bool
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
	Mu         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{RoomId: make(map[int]map[*Client]bool), Clients: make(map[*Client]bool), Broadcast: make(chan Message), Register: make(chan *Client), Unregister: make(chan *Client)}
}
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Mu.Lock()

			if _, ok := h.RoomId[client.RoomId]; ok {

			} else {
				h.RoomId[client.RoomId] = make(map[*Client]bool)

			}
			h.RoomId[client.RoomId][client] = true

			h.Mu.Unlock()
		case client := <-h.Unregister:
			h.Mu.Lock()
			if _, ok := h.RoomId[client.RoomId]; ok {
				delete(h.RoomId[client.RoomId], client)
				if len(h.RoomId[client.RoomId]) == 0 {

					delete(h.RoomId, client.RoomId)
				}
				close(client.Send)
			}
			h.Mu.Unlock()
		case message := <-h.Broadcast:
			h.Mu.Lock()
			if roomclients, ok := h.RoomId[message.RoomId]; ok {

				for Client := range roomclients {
					select {
					case Client.Send <- message.Data:
					default:
						close(Client.Send)
						delete(roomclients, Client)

					}

				}

			}
			if len(h.RoomId[message.RoomId]) == 0 {
				delete(h.RoomId, message.RoomId)
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
