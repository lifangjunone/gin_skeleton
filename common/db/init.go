package db

import (
	"gin_skeleton/common/db/mongo_impl"
	"gin_skeleton/common/db/mysql_impl"
)

type BaseDb interface{}

var (
	DB  *mysql_impl.MysqlDb
	MDB *mongo_impl.MongoDb
)

func InitMysqlDB() {
	DB = mysql_impl.InitMysqlDb()
}

func InitMongoDB() {
	MDB = mongo_impl.InitMongoDb()
}
