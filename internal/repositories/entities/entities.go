package entities

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
	UserTag   string    `json:"usertag" db:"usertag"`
	Name      string    `json:"name" db:"name"`
	Password  string    `db:"password"`
	Email     string    `db:"email"`
	CreatedAt time.Time `db:"ca"`
}
type Room struct {
	Id        int       `json:"idroom" db:"id"`
	User1_Id  int       `json:"user1_id" db:"user1_id"`
	User2_Id  int       `json:"user2_id" db:"user2_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
