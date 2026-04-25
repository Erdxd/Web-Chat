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
	ServiceU      service.UserService
	Hub           *Hub
	templates     *template.Template
	JwtMiddleware *middleware.JwtM
}

var msg struct {
	Text string `json:"text"`
}

func NewChatHandler(s *service.ServiceMessage, h *Hub, templates *template.Template, JwtMiddleware *middleware.JwtM, UserService service.UserService) *ChatHandler {
	return &ChatHandler{serviceM: *s, Hub: h, templates: templates, JwtMiddleware: JwtMiddleware, ServiceU: UserService}
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
			name, err := C.ServiceU.GetNameById(msg.UserId)
			HistoryMessage := dto.DtoMessage{
				Type:      "message",
				Id:        msg.Id,
				UserId:    msg.UserId,
				RoomId:    roomIdInt,
				CreatedAt: msg.CreatedAt,
				Content:   msg.Content,
				Name:      name,
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
		err = json.Unmarshal(payload, &msg)
		if err != nil {
			log.Println(err)
			return
		}

		MessageSave := model.Message{
			UserId:    Claims.User_id,
			RoomId:    roomIdInt,
			CreatedAt: time.Now(),
			Content:   msg.Text,
		}
		MessageId, err := C.serviceM.Save(MessageSave, Claims.User_id)
		name, err := C.ServiceU.GetNameById(MessageSave.UserId)

		if err != nil {
			http.Error(w, "Something is wrong", 500)
		}
		MessageDto := dto.DtoMessage{
			Type:      "message",
			Name:      name,
			Id:        MessageId,
			UserId:    Claims.User_id,
			RoomId:    roomIdInt,
			CreatedAt: time.Now(),
			Content:   msg.Text,
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
func (C *ChatHandler) FindUserByUserTag(w http.ResponseWriter, r *http.Request) {
	UserTag := r.FormValue("usertag")
	User, err := C.ServiceU.FindUserByUserTag(UserTag)
	if err != nil {
		http.Error(w, "Something is wrong", 500)
		return
	}
	C.templates.ExecuteTemplate(w, "searchuser.html", User)

}
