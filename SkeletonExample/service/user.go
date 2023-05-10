package service

import (
	"gin_skeleton/SkeletonExample/models"
	"gin_skeleton/SkeletonExample/service/impl"
)

var (
	UserSvc UserService
)

type UserService interface {
	GetUserById(id int) *models.User
}

func init() {
	UserSvc = impl.NewUserServiceImpl()
}
