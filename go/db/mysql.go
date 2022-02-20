package db

import (
	"go-web-demo/go/db/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SqlSession *gorm.DB

func InitMySql() (err error) {
	SqlSession, err = gorm.Open(mysql.Open(conf.GetDBConnectUrl()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := SqlSession.DB()
	sqlDB.SetMaxIdleConns(conf.GetMaxIdleConnCount())
	sqlDB.SetMaxOpenConns(conf.GetMaxOpenConnCount())
	sqlDB.SetConnMaxLifetime(conf.GetConnMaxLifetime())
	return err
}

func NewSqlSession() *gorm.DB {
	return SqlSession
}

func Close() {

}
