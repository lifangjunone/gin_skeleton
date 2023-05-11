package impl

import (
	"gin_skeleton/SkeletonExample/models"
	"gin_skeleton/common/db"
)

type DaoUserImpl struct{}

func NewDaoUser() *DaoUserImpl {
	return &DaoUserImpl{}
}

func (d *DaoUserImpl) Insert(u *models.User) {
	db.DB.Save(u)
}

func (d *DaoUserImpl) QueryById(user *models.User, id int) {
	db.DB.QueryById(user, id)
}
