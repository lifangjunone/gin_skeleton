package db

import (
	"gin_skeleton/common/db/mysql_impl"
)

type BaseDb interface{}

var (
	DB *mysql_impl.MysqlDb
)

func InitDB() {
	DB = mysql_impl.InitMysqlDb()
}
