package service

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository"
	"errors"
	"log"
	"time"
)

type RoomService struct {
	RoomRepo repository.Room
	UserRepo repository.User
}

func NewRoomService(RoomRepo repository.Room, UserRepo repository.User) *RoomService {
	return &RoomService{RoomRepo: RoomRepo, UserRepo: UserRepo}
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
func (RS *RoomService) GetAllPrivateChats(Userid1 int) ([]model.ChatWithName, error) {
	var Name string
	DataAboutRoom, err := RS.RoomRepo.GetAllPrivateChats(Userid1)
	if err != nil {
		errors.New("Something is wrong")
		log.Println(err)
	}
	var chats []model.ChatWithName
	for _, chat := range DataAboutRoom {
		if chat.UserId1 == Userid1 {
			Name, err = RS.UserRepo.GetNameUserById(chat.UserId2)
		} else {
			Name, err = RS.UserRepo.GetNameUserById(chat.UserId1)
		}
		if err != nil {
			errors.New("Something is wrong")
		}
		chats = append(chats, model.ChatWithName{Name: Name, Id: chat.ID})

	}
	return chats, err
}
