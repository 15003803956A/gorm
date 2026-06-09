package main

import (
	"fmt"
	"gorm/models"
	"gorm/utils"
)

func Delete() {
	//方式1
	/*var user = models.UserModel{Id: 11}
	utils.DB.Delete(&user)*/

	//方式2
	utils.DB.Delete(&models.UserModel{}, "id=?", 104)
	//utils.DB.Delete(&models.UserModel{}, "id=?", "8")
	//utils.DB.Delete(&models.UserModel{}, "name = ?", "兴达1")

	//批量删除
	utils.DB.Delete(&models.UserModel{}, []int{5, 109, 101})
}
func RuanDelete() {
	//查找软删除的数据
	var user models.UserModel
	utils.DB.Unscoped().Take(&user, "id=?", 104)
	fmt.Println(user)
	//真正删除
	utils.DB.Unscoped().Delete(&models.UserModel{}, "id=?", 104)
}
func main() {
	utils.LogInit()
	utils.DataBaseInit()
	//Delete()
	RuanDelete()

}
