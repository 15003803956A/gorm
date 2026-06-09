package main

import (
	"fmt"
	"gorm/models"
	"gorm/utils"

	"go.uber.org/zap"
)

/*func main() {
	//日志
	utils.LogInit()
	//数据库
	utils.DataBaseInit()
	//一次插入
	err := utils.DB.Create(&models.UserModel{
		Name:      "",
		Age:       0,
		Email:     "2675723732@qq.com",
		CreatedAt: time.Time{},
	}).Error
	if err != nil {
		zap.L().Error(err.Error())
	}
}*/
/*
func main() {
	//日志
	utils.LogInit()
	//数据库
	utils.DataBaseInit()
	//回填插入
	user := models.UserModel{
		Name: "兴达",
		Age:  18,
	}
	err := utils.DB.Create(&user).Error
	if err != nil {
		zap.L().Error(err.Error())
	}
	//被添加到表里的数据会被回填到user里
	fmt.Println(user)
}*/

/*func main() {
	//日志
	utils.LogInit()
	//数据库
	utils.DataBaseInit()
	//回填插入
	user := models.UserModel{
		Name: "兴达4",
		Age:  18,
	}
	err := utils.DB.Create(&user).Error
	if err != nil {
		zap.L().Error(err.Error())
	}
	//被添加到表里的数据会被回填到user里
	fmt.Println(user)
}*/

func main() {
	//日志
	utils.LogInit()
	//数据库
	utils.DataBaseInit()
	//回填插入
	user := []models.UserModel{{
		Name:  "小米",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "小小",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "小华",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "小伟",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "小欧",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "小剖",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "小薇",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "逍客",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "原始人",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "逍",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "葬帝",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "停产人",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "显示到人",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "神医",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "汉克",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "小贱贱",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "朱啸天",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "师弟",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "叶辰",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "黑袍",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "三尺青锋",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "天下",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "和人",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "不敢",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "纱",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "垃圾桶",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "之王",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "知网",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "一只小猪",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "猪舍",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "鸡舍",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "激素",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "糖果",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "石头",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "消息",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "短信",
		Age:   18,
		Email: "2675787878@qq.com",
	}, {
		Name:  "小强",
		Age:   18,
		Email: "2675787878@qq.com",
	}}
	err := utils.DB.Create(&user).Error
	if err != nil {
		zap.L().Error(err.Error())
	}
	//被添加到表里的数据会被回填到user里
	fmt.Println(user)
}
