package dao

import (
	"gin_skeleton/SkeletonExample/dao/impl"
	"gin_skeleton/SkeletonExample/ioc"
	"gin_skeleton/SkeletonExample/models"
)

type DaoUser interface {
	Insert(u *models.User)
	QueryById(user *models.User, id int)
}

func init() {
	ioc.DaoIocObj.Registry("daoUser", impl.NewDaoUser())
}
