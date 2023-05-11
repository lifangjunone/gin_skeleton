package impl

import (
	"gin_skeleton/SkeletonExample/models"
)

type UserServiceImpl struct {
	Name string
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{Name: "NewUserServiceImpl"}
}

func (u *UserServiceImpl) GetUserById(id int) *models.User {
	return &models.User{
		Id:       id,
		Username: "张三",
		Password: "123456",
	}
}
