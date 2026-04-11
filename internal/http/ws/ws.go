package http

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/service"
	"Web-Chat/internal/http/dto"
	"log"
	"net/http"
	"strconv"
	"text/template"
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
	service   service.ServiceMessage
	Hub       *Hub
	templates *template.Template
}

func NewChatHandler(s *service.ServiceMessage, h *Hub, templates *template.Template) *ChatHandler {
	return &ChatHandler{service: *s, Hub: h, templates: templates}
}
func (C *ChatHandler) OpenPipe(w http.ResponseWriter, r *http.Request) {
	roomId := r.URL.Query().Get("room")
	roomIdInt, err := strconv.Atoi(roomId)
	if err != nil {
		http.Error(w, "Cant parse your URl", 500)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &Client{
		Conn:   conn,
		Send:   make(chan []byte, 256),
		RoomId: roomIdInt,
	}
	C.Hub.Register <- client
	go client.WritePump()
	defer func() {
		C.Hub.Unregister <- client
		conn.Close()

		if err != nil {
			log.Println(err)
			http.Error(w, "1", 400)
		}
	}()
	message, err := C.service.CheckMessage(roomIdInt)
	if err != nil {
		log.Println(err)
		return
	} else {
		for _, msg := range message {

			client.Send <- []byte(msg.Content)
		}
	}
	for {

		_, payload, err := conn.ReadMessage()
		if err != nil {
			break
		}

		MessageSave := dto.DtoMessage{
			Id:        1,
			UserId:    1,
			RoomId:    roomIdInt,
			CreatedAt: time.Now(),
			Content:   string(payload),
		}
		log.Println(MessageSave)
		err = C.service.Save(model.Message(MessageSave), 1)

		BroadCast := Message{
			Data:   payload,
			RoomId: roomIdInt,
		}
		C.Hub.Broadcast <- BroadCast

	}

}
