package service

import (
	"gin_skeleton/SkeletonExample/ioc"
	"gin_skeleton/SkeletonExample/models"
	"gin_skeleton/SkeletonExample/service/impl"
)

type UserService interface {
	GetUserById(user *models.User, id interface{})
	CreateUser(user *models.User)
	QueryByField(map[string]interface{}) (users *[]models.User)
}

func init() {
	ioc.SvcIocObj.Registry("UserService", impl.NewUserServiceImpl())
}
