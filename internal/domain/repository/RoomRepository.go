package repository

import (
	"Web-Chat/internal/domain/model"
	"time"
)

type Room interface {
	DeleteRoom(RoomId int) error
	FindRoomIdByUsersId(UserId1, UserId2 int) (int, error)
	FindUsersByRoomId(RoomId int) (int, int, error)
	CreateRoom(UserId1, UserId2 int, ca time.Time) (int, error)
	GetAllPrivateChats(UserId1 int) ([]model.PrivateChat, error)
}
