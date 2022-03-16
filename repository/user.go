package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/plancer16/mall_end/model"
)

type UserRepo struct {
	DB *gorm.DB
}

func (repo *UserRepo) Exist(user model.User) *model.User {
	var count int
	repo.DB.Find(&user).Where("nick_name=?", user.NickName).Count(&count)
	if count > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepo) Add(user model.User) (*model.User, error) {
	if exist := repo.Exist(user); exist != nil {
		return nil, fmt.Errorf("用户已经注册")
	}
	err := repo.DB.Create(&user).Error
	if err != nil {
		return nil, fmt.Errorf("用户注册失败")
	}
	return &user, nil
}