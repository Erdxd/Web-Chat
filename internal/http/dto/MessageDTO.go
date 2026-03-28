package dto

import "time"

type DtoMessage struct {
	Id        int64     `json:"idmessage"`
	UserId    int       `json:"useridmessage"`
	RoomId    int       `json:"roomidmessage"`
	CreatedAt time.Time `json:"CAmessage"`
	Content   string    `json:"message"`
}
