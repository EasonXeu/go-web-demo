package main

import (
	"go-web-demo/src/dal/dataobject"
	"go-web-demo/src/db"
)

func main() {
	//连接数据库
	err := db.InitMySql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	defer db.Close()
	//绑定模型
	db.SqlSession.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(&dataobject.User{})

	//注册路由
	engine := SetRouter()
	//启动端口为8085的项目
	engine.Run(":8081")
}
