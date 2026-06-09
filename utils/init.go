package utils

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 结构体字段名需要大写

var DB *gorm.DB

func DataBaseInit() {
	//数据库
	DSN := "root:123456@tcp(127.0.0.1:3306)/gorm_db_new?charset=utf8&parseTime=true"

	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
}
