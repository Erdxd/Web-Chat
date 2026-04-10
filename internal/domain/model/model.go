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
type Claims struct {
	User_id int
	Admin   bool

	jwt.RegisteredClaims
}
