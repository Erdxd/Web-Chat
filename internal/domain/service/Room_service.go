package service

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository"
	"time"
)

type RoomService struct {
	RoomRepo repository.Room
}

func NewRoomService(RoomRepo repository.Room) *RoomService {
	return &RoomService{RoomRepo: RoomRepo}
}
func (RS *RoomService) CreateNewChat(UserId1, UserId2 int, ca time.Time) (int, error) {
	return RS.RoomRepo.CreateRoom(UserId1, UserId2, ca)
}
func (RS *RoomService) DeleteRoom(RoomId int) error {
	return RS.RoomRepo.DeleteRoom(RoomId)
}
func (RS *RoomService) FindRoomIdByUsersId(UserId1 int, UserId2 int) (int, error) {
	return RS.RoomRepo.FindRoomIdByUsersId(UserId1, UserId2)
}
func (RS *RoomService) FindUsersByRoomId(RoomId int) (int, int, error) {
	return RS.RoomRepo.FindUsersByRoomId(RoomId)
}
func (RS *RoomService) GetAllPrivateChats(Userid1 int) ([]model.PrivateChat, error) {
	return RS.RoomRepo.GetAllPrivateChats(Userid1)
}
