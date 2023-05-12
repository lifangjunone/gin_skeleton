package impl

import (
	"gin_skeleton/SkeletonExample/models"
	"gin_skeleton/common/db"
	"strconv"
)

var colName = "user"

type DaoUserMongoImpl struct{}

func NewDaoMongoUser() *DaoUserMongoImpl {
	return &DaoUserMongoImpl{}
}

func (d *DaoUserMongoImpl) Insert(u *models.User) {
	db.MDB.InsertOne(colName, u)
}

func (d *DaoUserMongoImpl) QueryById(user *models.User, id int) {
	db.MDB.QueryById(colName, strconv.Itoa(id))
}
