package main

import (
	"fmt"
	"gorm/models"
	"gorm/utils"
)

//func main() {
//	utils.LogInit()
//	utils.DataBaseInit()
//
//	err := utils.DB.AutoMigrate(&models.Company{})
//	if err != nil {
//		zap.L().Error(err.Error())
//		return
//	}
//	err = utils.DB.AutoMigrate(&models.User{})
//	if err != nil {
//		zap.L().Error(err.Error())
//		return
//	}
//}

func init() {
	utils.LogInit()
	utils.DataBaseInit()
}
func belongsTo() {
	//c := models.Company{
	//	Name: "北邮有限公司",
	//}
	//u := models.User{
	//	Name:    "小明",
	//	Age:     18,
	//	Company: c,
	//}
	//err := utils.DB.Create(&u).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err := utils.DB.AutoMigrate(&models.Company{}, &models.User{})
	//if err != nil {
	//	zap.L().Error("belongsTo():", zap.Error(err))
	//	return
	//}
}

func hasOneInit() {
	err := utils.DB.AutoMigrate(&models.Girl{}, &models.Dog{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
func hasOne() {
	//g := models.Girl{GirlName: "淼1"}
	//utils.DB.Create(&g)
	//
	//d := models.Dog{
	//	Name:   "杨栋旭",
	//	GirlID: g.Id,
	//}
	//utils.DB.Create(&d)

	//g := models.Girl{}
	//utils.DB.Preload("Dog").Find(&g, "4")
	//fmt.Println(g)

	//utils.DB.Delete(&models.Girl{}, "4")
	hasOneInit()
}
func hasMany() {
	//err := utils.DB.AutoMigrate(&models.Girl{}, &models.Dag{})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
}
func main() {
	//belongsTo()
	hasOne()
	//hasMany()
}
