package models

//天狗被打上主人的标记
//type User struct {
//	Id        int    `gorm:"type:int;size:32;primary_key;AUTO_INCREMENT"`
//	Name      string `gorm:"type:varchar(255)"`
//	Age       int    `gorm:"type:int;size:32"`
//	CompanyID int    `gorm:"type:int;size:32"`
//	Company   Company
//}
//type Company struct {
//	Id   int    `gorm:"type:int;size:32;primary_key;AUTO_INCREMENT"`
//	Name string `gorm:"type:varchar(255)"`
//}

// 主人拥有舔狗has one
type Dog struct {
	Id     int `gorm:"primary_key;auto_increment"`
	Name   string
	GirlID int
}

type Girl struct {
	Id       int `gorm:"primary_key;auto_increment"`
	GirlName string
	Dog      Dog
}

//type Dag struct {
//	Id     int    `gorm:"type:int;size:32;primary_key;auto_increment"`
//	Name   string `gorm:"type:varchar(255);"`
//	GirlID int    `gorm:"type:int;size:32;"`
//}
//
//type Girl struct {
//	Id   int    `gorm:"type:int;size:32;primary_key;auto_increment"`
//	Name string `gorm:"type:varchar(255);"`
//	Dag  []Dag
//}
