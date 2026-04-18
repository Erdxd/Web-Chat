package dto

import "time"

type DtoMessage struct {
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Id        int64     `json:"id"`
	UserId    int       `json:"userid"`
	RoomId    int       `json:"roomid"`
	CreatedAt time.Time `json:"createdAt"`
	Content   string    `json:"content"`
}
