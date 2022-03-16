package service

import (
	"github.com/plancer16/mall_end/model"
	"github.com/plancer16/mall_end/repository"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	UserRepo repository.UserRepo
}

func (userService *UserService) Add(user model.User) (*model.User,error) {
	user.UserId = uuid.NewV4().String()
	if user.Password == "" {
		user.Password = "123456"
	}
	user.IsDeleted = false
	user.IsLocked = false
	return userService.UserRepo.Add(user)
}