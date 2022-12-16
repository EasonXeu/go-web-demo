package db

import (
	"go-web-demo/src/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SqlSession *gorm.DB

func InitMySql() (err error) {
	cfg, err := conf.LoadConfig()
	if err != nil {
		panic(err)
	}
	SqlSession, err = gorm.Open(mysql.Open(cfg.GetDBConnectUrl()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := SqlSession.DB()
	sqlDB.SetMaxIdleConns(cfg.GetMaxIdleConnCount())
	sqlDB.SetMaxOpenConns(cfg.GetMaxOpenConnCount())
	sqlDB.SetConnMaxLifetime(cfg.GetConnMaxLifetime())
	return err
}

func NewSqlSession() *gorm.DB {
	return SqlSession
}

func Close() {

}
