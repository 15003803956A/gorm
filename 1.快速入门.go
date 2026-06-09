package main

import (
	"fmt"
	"gorm/utils"
)

// import (
//
//	"context"
//	"gorm.io/driver/sqlite"
//	"gorm.io/gorm"
//
// )
//
//	type Product struct {
//		gorm.Model
//		Code  string
//		Price uint
//	}
//
//	func main() {
//		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
//		if err != nil {
//			panic("连接数据库失败")
//		}
//
//		ctx := context.Background()
//
//		// 自动建表
//		db.AutoMigrate(&Product{})
//
//		// 创建
//		err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})
//
//		// 查询
//		product, err := gorm.G[Product](db).Where("id = ?", 1).First(ctx) // 查找对应主键的产品
//		products, err := gorm.G[Product](db).Where("code = ?", "D42").Find(ctx) // 查找 code 为 D42 的所有产品
//
//		// 更新 - 将产品价格更新为 200
//		err = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
//		// 更新 - 更新多个字段
//		err = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, Product{Code: "D42", Price: 100})
//
//		// 删除 - 删除产品
//		err = gorm.G[Product](db).Where("id = ?", product.ID).Delete(ctx)
//	}
func main() {
	utils.DataBaseInit()
	var version string
	utils.DB.Raw("select version()").Scan(&version)
	fmt.Println(version)
}
