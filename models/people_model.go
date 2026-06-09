package models

type PeopleModel struct {
	Name     string `gorm:"type:varchar(255);unique"`
	Age      int    `gorm:"type:int(32)"`
	Email    string `gorm:"type:varchar(255)"`
	Birthday string `gorm:"type:datetime"`
}
