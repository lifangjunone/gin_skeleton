package mysql_impl

import (
	"fmt"
	"gin_skeleton/SkeletonExample/models"
	"gin_skeleton/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDb struct {
	Db *gorm.DB
}

func (g *MysqlDb) Save(o interface{}) {
	g.Db.Create(o)
}

func (g *MysqlDb) Update(o interface{}) {
	g.Db.Updates(o)
}

func (g *MysqlDb) QueryById(user *models.User, id interface{}) {
	g.Db.First(user, id)
}

func (g *MysqlDb) QueryByIds(users []*models.User, ids []int) {
	g.Db.Find(users, ids)
}

func (g *MysqlDb) Exec(sqlStr string, users []*models.User) {
	g.Db.Exec(sqlStr, users)
}

func InitMysqlDb() *MysqlDb {
	mysqlConf := global.ConfObj.Mysql
	sqlStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		mysqlConf.UserName, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.Database)
	gormDB, err := gorm.Open(mysql.Open(sqlStr), &gorm.Config{}) //配置项中预设了连接池 ConnPool
	if err != nil {
		fmt.Println("数据库连接出现了问题：", err)
		panic(err.Error())
	}
	return &MysqlDb{
		Db: gormDB,
	}
}
