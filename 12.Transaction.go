package main

import (
	"fmt"
	"gorm/models"
	"gorm/utils"

	"gorm.io/gorm"
)

func init() {
	utils.LogInit()
	utils.DataBaseInit()
	userModelInit()
}
func userModelInit() {
	err := utils.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 有两种使用事务的方式，本质上一样
// 第一种:
func transactionHalfAuto() {
	err := utils.DB.Transaction(func(tx *gorm.DB) error {
		//-100
		err := tx.Model(&models.UserModel{}).
			Debug().
			Where("id", 1).
			Update("age", gorm.Expr("age - ?", 100)).
			Error
		//err = errors.New("截断")
		if err != nil {
			tx.Rollback()
		}
		//+100
		err = tx.Model(&models.UserModel{}).
			Debug().
			Where("id", 2).
			Update("age", gorm.Expr("age + ?", 100)).
			Error
		if err != nil {
			tx.Rollback()
		}
		//提交
		tx.Commit()
		return err
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func transactionManual() {
	tx := utils.DB.Begin()
	err := tx.Model(&models.UserModel{}).
		Where("id", 1).
		Update("age", gorm.Expr("age + ?", 100)).
		Error
	//err = errors.New("截断")
	if err != nil {
		tx.Rollback()
	}
	err = tx.Model(&models.UserModel{}).
		Where("id", 2).
		Update("age", gorm.Expr("age - ?", 100)).
		Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
}

func transactionManualToName() {
	tx := utils.DB.Begin()
	err := tx.Model(&models.UserModel{}).
		Where("id", 6).
		Update("name", "杨栋旭1").
		Error
	//err = errors.New("截断")
	if err != nil {
		tx.Rollback()
	}
	err = tx.Model(&models.UserModel{}).
		Where("id", 7).
		Update("name", "王鹏召").
		Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()
}
func main() {
	//transactionManual()
	//transactionHalfAuto()
	transactionManualToName()
}
