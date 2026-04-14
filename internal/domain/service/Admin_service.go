package service

import (
	"Web-Chat/internal/domain/model"
	"Web-Chat/internal/domain/repository"
)

type AdminSerivce struct {
	AdminRepo repository.Admin
}

func NewAdminService(AdminRepo repository.Admin) *AdminSerivce {
	return &AdminSerivce{AdminRepo: AdminRepo}
}
func (AS *AdminSerivce) CheckAllUsers() ([]model.User, error) {
	return AS.AdminRepo.CheckAllUsers()
}
func (AS *AdminSerivce) FoundUserByUserId(UserId int) (model.User, error) {
	return AS.AdminRepo.FoundUserByUserId(UserId)
}
func (AS *AdminSerivce) DeleteUser(UserId int) error {
	return AS.AdminRepo.DeleteUser(UserId)
}
