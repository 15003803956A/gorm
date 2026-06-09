package main

import (
	"fmt"
	"gorm/models"
	"gorm/utils"
)

func query() {
	//查询所有
	//var userlist []models.UserModel
	//utils.DB.Find(&userlist)
	//fmt.Println(userlist)

	//条件查询
	//utils.DB.Find(&userlist, "name = ?", "光头强")
	//fmt.Println(userlist)

	//查询单个
	var user models.UserModel
	//utils.DB.Take(&user)
	//fmt.Println(user)

	//查询单个,带条件
	//utils.DB.Take(&user, "name = ?", "兴达")
	//fmt.Println(user)

	//查询单个,带条件,如果有多条数据，只返回第一条
	//utils.DB.Take(&user, "age > ?", "18")
	//fmt.Println(user)

	//根据主键查询
	//utils.DB.Take(&user, 6)
	//fmt.Println(user)

	//根据主键查询,如果主键user.Id本来就有有值，在查询的时候会根据这个主键Id查，但这只限于主键
	//user.Id = 8
	//utils.DB.Take(&user)
	//fmt.Println(user)

	//按主键id升序，取第一条满足条件的数据
	utils.DB.First(&user, "age = ?", "200")
	fmt.Println(user)

	//按主键id升序，取最后一条满足条件的数据
	//utils.DB.Last(&user, "age = ?", "200")
	//fmt.Println(user)

	//下面属于错示例，user.Id=8，同时又指定查询主键=8，此时它会查询id=8 and id=6的数据
	//user.Id = 8
	//utils.DB.Take(&user, 6)
	//fmt.Println(user)

	//当没有查到数据，会报错：record not found，需要处理错误
	//err := utils.DB.Take(&user, 100).Error
	//if err != nil {
	//	zap.L().Error("查询为空:" + err.Error())
	//}
	//fmt.Println(user)

	//.Debug或打印出生成的SQL语句
	//utils.DB.Debug().Find(&user, 6)
	//fmt.Println(user)
}

func main() {
	utils.LogInit()
	utils.DataBaseInit()
	query()
}
