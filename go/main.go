package main

import (
	"go-web-demo/go/dal/dataobject"
	"go-web-demo/go/db"
	"go-web-demo/go/routes"
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
	engine := routes.SetRouter()
	//启动端口为8085的项目
	engine.Run(":8081")
}
