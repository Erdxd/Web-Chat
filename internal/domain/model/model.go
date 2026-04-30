package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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
type Room struct {
	Id        int
	User1_Id  int
	User2_Id  int
	CreatedAt time.Time
}
type Claims struct {
	User_id int
	Admin   bool

	jwt.RegisteredClaims
}
type UserView struct {
	Email   string
	Name    string
	Ca      time.Time
	Usertag string
}
type UserSerchResult struct {
	UserTag string
	Name    string
	UserId  int
}
type PrivateChat struct {
	UserId1 int
	UserId2 int
	ID      int
}
type ChatWithName struct {
	Name   string
	Id     int
	Online bool
}
