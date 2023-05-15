package impl

import (
	"gin_skeleton/SkeletonExample/dao"
	"gin_skeleton/SkeletonExample/ioc"
	"gin_skeleton/SkeletonExample/models"
)

var (
	userDao = ioc.DaoIocObj.GetSvc("daoUser").(dao.DaoUser)
)

type UserServiceImpl struct {
	Name string
}

func NewUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{Name: "NewUserServiceImpl"}
}

func (u *UserServiceImpl) GetUserById(user *models.User, id interface{}) {
	userDao.QueryById(user, id)
}

func (u *UserServiceImpl) CreateUser(user *models.User) {
	userDao.Insert(user)
}

func (u *UserServiceImpl) QueryByField(jsonParam map[string]interface{}) *[]models.User {
	return userDao.QueryByField(jsonParam)
}
