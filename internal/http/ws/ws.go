package http

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/service"
	"Web-Chat/internal/http/dto"
	"Web-Chat/internal/http/middleware"
	"encoding/json"
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
	serviceM      service.ServiceMessage
	Hub           *Hub
	templates     *template.Template
	JwtMiddleware *middleware.JwtM
}

func NewChatHandler(s *service.ServiceMessage, h *Hub, templates *template.Template, JwtMiddleware *middleware.JwtM) *ChatHandler {
	return &ChatHandler{serviceM: *s, Hub: h, templates: templates, JwtMiddleware: JwtMiddleware}
}
func (C *ChatHandler) OpenPipe(w http.ResponseWriter, r *http.Request) {
	roomId := r.URL.Query().Get("room")
	roomIdInt, err := strconv.Atoi(roomId)
	if err != nil {
		http.Error(w, "Cant parse your URl", 500)
		return
	}
	Claims, err := C.JwtMiddleware.GetDataFromJwt(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", 401)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	log.Println("websokets1")

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

		}
	}()
	log.Println("websokets2")

	message, err := C.serviceM.CheckMessage(roomIdInt)
	if err != nil {
		log.Println(err)
		return
	} else {
		for _, msg := range message {
			HistoryMessage := dto.DtoMessage{
				Type:      "message",
				Id:        msg.Id,
				UserId:    msg.UserId,
				RoomId:    roomIdInt,
				CreatedAt: msg.CreatedAt,
				Content:   msg.Content,
			}
			JsonData, err := json.Marshal(HistoryMessage)
			if err != nil {
				http.Error(w, "Cant Create JSON", 500)
			}
			client.Send <- JsonData
		}
	}
	for {

		_, payload, err := conn.ReadMessage()
		if err != nil {
			break
		}

		MessageSave := model.Message{

			UserId:    Claims.User_id,
			RoomId:    roomIdInt,
			CreatedAt: time.Now(),
			Content:   string(payload),
		}
		MessageId, err := C.serviceM.Save(MessageSave, Claims.User_id)
		if err != nil {
			http.Error(w, "Something is wrong", 500)
		}
		MessageDto := dto.DtoMessage{
			Type:      "message",
			Id:        MessageId,
			UserId:    Claims.User_id,
			RoomId:    roomIdInt,
			CreatedAt: time.Now(),
			Content:   string(payload),
		}

		JsonData, err := json.Marshal(MessageDto)
		if err != nil {
			http.Error(w, "Cant Create JSON", 500)
		}
		C.Hub.Broadcast <- Message{
			Data:   JsonData,
			RoomId: roomIdInt,
		}

	}

}
func (C *ChatHandler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {

		Claims, err := C.JwtMiddleware.GetDataFromJwt(w, r)
		if err != nil {
			http.Error(w, "Unauthorized", 401)
			return
		}
		RoomId, err := strconv.Atoi(r.URL.Query().Get("room"))
		if err != nil {
			http.Error(w, "Cant find room", 401)
			return
		}
		IdMessage, err := strconv.Atoi(r.URL.Query().Get("messageid"))
		if err != nil {
			http.Error(w, "Cant find your message", 401)
			return
		}
		err = C.serviceM.DeleteMessage(Claims.User_id, RoomId, int64(IdMessage))
		if err != nil {
			http.Error(w, "Cant delete this message", 400)
			return
		}
		DeleteMessage := dto.DtoMessage{
			Type:   "delete",
			Id:     int64(IdMessage),
			UserId: Claims.User_id,
			RoomId: RoomId,
		}
		JsonData, err := json.Marshal(DeleteMessage)
		if err != nil {
			http.Error(w, "Cant parse JSON", 401)
			return
		}
		C.Hub.Broadcast <- Message{
			Data:   JsonData,
			RoomId: RoomId,
		}
		w.WriteHeader(http.StatusOK)

	}

}
