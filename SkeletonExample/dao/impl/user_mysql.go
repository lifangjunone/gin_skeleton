package impl

import (
	"gin_skeleton/SkeletonExample/models"
	"gin_skeleton/common/db"
)

type DaoUserMysqlImpl struct{}

func NewDaoMysqlUser() *DaoUserMysqlImpl {
	return &DaoUserMysqlImpl{}
}

func (d *DaoUserMysqlImpl) Insert(u *models.User) {
	db.DB.Save(u)
}

func (d *DaoUserMysqlImpl) QueryById(user *models.User, id interface{}) {
	db.DB.QueryById(user, id)
}
