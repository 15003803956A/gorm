package main

import (
	"fmt"
	"gorm/models"
	"gorm/utils"
)

func HighQuery() {
	//1.直接写条件
	//var user models.UserModel
	//utils.DB.Where("age = ?", 0).Take(&user)
	//utils.DB.Where("name = ?", "xxx").Take(&user)
	//utils.DB.Debug().Where("name != ?", nil).Take(&user)
	//fmt.Println(user)

	//2.传结构体,不会传零值
	//var user models.UserModel
	//utils.DB.Debug().Where(models.UserModel{Age: 0, Name: "xxx"}).Take(&user)
	//fmt.Println(user)

	//2.传结构体,不会传零值
	//var users []models.UserModel
	//utils.DB.Debug().Where(models.UserModel{Age: 18}).Find(&users)
	//fmt.Println(users)

	//3.传map，会传零值
	//var user models.UserModel
	//utils.DB.Debug().Where(map[string]any{"age": 0}).Take(&user)
	//fmt.Println(user)

	//4.where嵌套
	//var user models.UserModel
	//query := utils.DB.Debug().Where(models.UserModel{Name: "xxx"})
	//utils.DB.Where(query).Take(&user)
	//fmt.Println(user)

	//5.or
	//var user []models.UserModel
	//utils.DB.Or("name = ?", "xxx").Or("name = ?", "rrr").Find(&user)
	//fmt.Println(user)

	//6.not
	//	var users []models.UserModel
	//	utils.DB.Debug().Not("age < ?", 18).Not("name = ?", "xxx").Find(&users)
	//	fmt.Println(users)

	//7.order
	var users []models.UserModel
	utils.DB.Debug().Order("age desc").Order("id asc").Find(&users)
	Print(users)
}
func Print(users []models.UserModel) {
	for _, user := range users {
		fmt.Println(user)
	}
}

// Model(models.UserModel{})就是在显式指定数据来源。
// Model(&models.UserModel{})里的参数，当增删改时加&，查找不加&
func Scan() {
	//1.scan取全部的name
	//var nameList []string
	//utils.DB.Debug().Model(models.UserModel{}).Select("name").Scan(&nameList)
	//fmt.Println(nameList)

	//2.也可以使用Pluck
	//var nameList []string
	//utils.DB.Debug().Model(models.UserModel{}).Pluck("name", &nameList)
	//fmt.Println(nameList)

	//3.可以新建一个字段比较少的结构体，只接受自己想要的字段
	//type userLow struct {
	//	Name string
	//	Age  int
	//}
	//var userLowList []userLow
	//utils.DB.Debug().Model(&models.UserModel{}).Scan(&userLowList)
	//fmt.Println(userLowList)

	//4.`gorm:"column:name"`使用column可以指定db列名
	//type userLow struct {
	//	Username string `gorm:"column:name"`
	//	Agevalue int    `gorm:"column:age"`
	//}
	//var userLowList []userLow
	//utils.DB.Debug().Model(models.UserModel{}).Scan(&userLowList)
	//fmt.Println(userLowList)

	//5.group分组
	//type userLowGroup struct {
	//	Age   int
	//	Count int
	//}
	//var userLowGroupList []userLowGroup
	//utils.DB.Debug().Model(models.UserModel{}).Group("age").Select("age,count(*) as count").Scan(&userLowGroupList)
	//fmt.Println(userLowGroupList)

	//6.去重
	//var nameList []string
	//utils.DB.Debug().Model(models.UserModel{}).Select("name").Distinct().Scan(&nameList)
	//fmt.Println(nameList)

	//
	//type User struct {
	//	Age   int
	//	Count int
	//}
	//var users []User
	//utils.DB.Debug().Model(models.UserModel{}).Group("age").Select("age,count(*) as count").Scan(&users)
	//fmt.Println(users)
}

func Offset() {
	//offset
	var users []models.UserModel
	// 第一页
	utils.DB.Limit(10).Offset(0).Find(&users)
	fmt.Println(users)
	// 第二页
	utils.DB.Limit(10).Offset(10).Find(&users)
	fmt.Println(users)
	// 第三页
	utils.DB.Limit(10).Offset(20).Find(&users)
	fmt.Println(users)
	// 第n页
	//var n = (页码-1)*一页条数
	//utils.DB.Limit(10).Offset(n).Find(&users)
}

// Scopes(funcs ...func(*DB) *DB)，因此参数类型可以有多个，但必须是func(*DB) *DB
func Scope() {
	//Age18
	//var users []models.UserModel
	//utils.DB.Scopes(Age18).Find(&users)
	//fmt.Println(users)

	//GetByName
	//var users []models.UserModel
	//utils.DB.Debug().Scopes(GetByName).Find(&users)
	//fmt.Println(users)

	//可以写多个参数
	//var users []models.UserModel
	//utils.DB.Debug().Scopes(GetByName, OmitEmail).Find(&users)
	//fmt.Println(users)
}

//func Age18(tx *gorm.DB) *gorm.DB {
//	return tx.Where("age >= ?", 18)
//}
//func GetByName(tx *gorm.DB) *gorm.DB {
//	return tx.Where("name = ?", "xxx")
//}
//func OmitEmail(tx *gorm.DB) *gorm.DB {
//	return tx.Omit("email")
//}

func scopePlus() {
	//带参数的scope函数,写法:写一个函数，以func(tx *gorm.DB) *gorm.DB作为返回值类型
	//var users []models.UserModel
	//utils.DB.Scopes(NameIn("xxx", "rrr", "r")).Find(&users)
	//fmt.Println(users)
}

//	func NameIn(nameList ...string) func(tx *gorm.DB) *gorm.DB {
//		return func(tx *gorm.DB) *gorm.DB {
//			return tx.Where("name in  ?", nameList)
//		}
//	}
func Row() {
	//1.Row原生SQL
	//var users []models.UserModel
	//utils.DB.Raw("select * from user_models where name in ?", []string{"xxx", "rrr", "r"}).Scan(&users)
	//fmt.Println(users)

	//2.Exec原生SQL
	utils.DB.Exec("update user_models set name =? where name = ?", "lll", "xxx")
}
func main() {
	utils.LogInit()
	utils.DataBaseInit()
	//HighQuery()
	//Scan()
	//Offset()
	//Scope()
	//scopePlus()
	Row()
}
