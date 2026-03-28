package model

import "time"

type Message struct {
	Id        int64
	UserId    int
	RoomId    int
	CreatedAt time.Time
	Content   string
}
type User struct {
	UserId    int
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
}
