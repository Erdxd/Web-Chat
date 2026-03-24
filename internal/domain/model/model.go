package model

import "time"

type Message struct {
	Id        int64     `json:"idmessage" db:"idmessage"`
	UserId    int       `json:"useridmessage" db:"useridmessage"`
	RoomId    int       `json:"roomidmessage" db:"roomidmessage"`
	CreatedAt time.Time `json:"CAmessage" db:"CAmessage"`
	Content   string    `json:"message" db:"message"`
}
type User struct {
	UserId    int       `json:"userid" db:"userid"`
	Name      string    `db:"name"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"CA"`
}
