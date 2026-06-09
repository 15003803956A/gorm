package main

import (
	"fmt"
	"gorm/models"
	"gorm/utils"
	"time"

	"gorm.io/gorm"
)

/*
1.Save，有主键记录就是更新，并且可以更新零值，否则就是创建
2.Update，可以更新零值，必须要有条件
3. UpdateColumn，可以更新零值，不会走更新的Hook
4.Updates，如果是结构体，则更新非零值，map可以更新零值
*/

func save() {
	//id=1存在，更新
	//utils.DB.Save(&models.UserModel{
	//	Id:        1,
	//	Name:      "xx",
	//	Age:       18,
	//	Email:     "6373637176@163.com",
	//	CreatedAt: time.Now(),
	//})

	//id=100不存在，插入
	utils.DB.Save(&models.UserModel{
		Id:        100,
		Name:      "xxx",
		Age:       18,
		Email:     "6373637176@163.com",
		CreatedAt: time.Now(),
	})
}

// 更新操作会回填被成功更新的值
func update() {
	var user = models.UserModel{Id: 1}
	utils.DB.Model(&user).Update("name", "zzz")
	fmt.Println(user)
}

// 和update唯一的区别就是updateClounm会跳过钩子函数，tx.Statement.SkipHooks = true
func updateClounm() {
	var user = models.UserModel{Id: 1}
	utils.DB.Model(&user).UpdateColumn("age", "0")
	fmt.Println(user)
}

func updates() {
	//正常使用
	/*var user = models.UserModel{Id: 1}
	utils.DB.Model(&user).Updates(models.UserModel{
		Name: "杨栋旭",
		Age:  10,
	})
	fmt.Println(user)*/

	//如果值是零值,不会更新
	/*var user = models.UserModel{Id: 1}
	utils.DB.Model(&user).Updates(models.UserModel{
		Name: "杨栋xx旭",
		Age:  0,
	})
	fmt.Println(user)*/

	//使用map进行更新，可以更新零值
	var user = models.UserModel{Id: 1}
	utils.DB.Model(&user).Updates(map[string]any{
		"name": "",
		"age":  0,
	})
	fmt.Println(user)
}

// Expr:这部分值不要当普通数据处理，而是直接作为SQL表达式执行.通常用于获取原字段的数据，例如年龄加1
func expr() {
	var user = models.UserModel{Id: 1}
	utils.DB.Model(&user).Update("age", gorm.Expr("age + ?", 1))
	fmt.Println(user)
}
func main() {
	utils.LogInit()
	utils.DataBaseInit()
	//save()
	//update()
	//updateClounm()
	//updates()
	expr()
}
