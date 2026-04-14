package dto

import "time"

type User struct {
	UserId    int
	Name      string
	Password  string
	Email     string
	CreatedAt time.Time
}
