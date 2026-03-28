package http

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/service"
	"Web-Chat/internal/http/dto"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type ChatHandler struct {
	service service.ServiceMessage
	Hub     *Hub
}

func NewChatHandler(s *service.ServiceMessage, h *Hub) *ChatHandler {
	return &ChatHandler{service: *s, Hub: h}
}
func (C *ChatHandler) OpenPipe(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	client := &Client{
		Conn: conn,
		Send: make(chan []byte, 256),
	}
	C.Hub.Register <- client
	go client.WritePump()
	defer func() {
		C.Hub.Unregister <- client
		conn.Close()
	}()
	for {
		_, payload, err := conn.ReadMessage()
		if err != nil {
			break
		}
		MessageSave := dto.DtoMessage{
			Id:        1,
			UserId:    1,
			RoomId:    1,
			CreatedAt: time.Now(),
			Content:   string(payload),
		}
		log.Println(MessageSave)
		err = C.service.Save(model.Message(MessageSave), 1)
		C.Hub.Broadcast <- payload

	}
}
