package main

import (
	"gorm/utils"

	"go.uber.org/zap"
)

// 结构体字段名需要大写，根据字段名自动和数据库列名对应在一起
type User struct {
	Id    int
	Name  string
	Email string
	Age   int
}

//	func main() {
//		//日志
//		utils.LogInit()
//		//连接数据库
//		DSN := "root:123456@tcp(127.0.0.1:3306)/gorm_db_new"
//		db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
//		if err != nil {
//			zap.L().Error("数据库连接失败:" + err.Error())
//			return
//		}
//		//查询所有的User
//		var userlist []User
//		db.Find(&userlist)
//		fmt.Println(userlist)
//	}
func main() {
	utils.LogInit()
	zap.L().Error("assasa")
}
