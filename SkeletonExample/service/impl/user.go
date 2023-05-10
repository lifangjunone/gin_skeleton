package impl

import (
	"gin_skeleton/SkeletonExample/models"
)

type UserServiceImpl struct{}

func (u *UserServiceImpl) GetUserById(id int) *models.User {
	return &models.User{
		Id:       id,
		Username: "张三",
		Password: "123456",
	}
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}
