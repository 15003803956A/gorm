package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UserModel struct {
	Id        int       `json:"id"`
	Name      string    `json:"name" gorm:"not null;unique" `
	Age       int       `json:"age" gorm:"default:0"`
	CreatedAt time.Time `json:"createdat" gorm:"type:datetime(0)"`
}

// 返回nil,插入成功,可以在这里对要插入的数据进行修改
//func (u *UserModel) BeforeCreate(db *gorm.DB) error {
//	u.Name = u.Name + "Plus"
//	return nil
//}

//返回err,不会插入
/*func (u *UserModel) BeforeCreate(db *gorm.DB) error {
	fmt.Println("钩子函数触发了")
	return errors.New("xxx")
}*/

// 返回nil,允许更新,可以在这里对要更新的数据进行修改
/*func (u *UserModel) BeforeUpdate(db *gorm.DB) error {
	u.Name = u.Name + "Plus"
	return errors.New("as")
}*/

// 返回nil,允许更新,可以在这里对要更新的数据进行修改
/*func (u *UserModel) BeforeUpdate(db *gorm.DB) error {
	u.Name = u.Name + "Plus"
	return nil
}*/

/*func (u *UserModel) BeforeDelete(db *gorm.DB) error {
	fmt.Println("钩子函数BeforeDelete触发了")
	return nil
}*/
/*func (u *UserModel) AfterDelete(db *gorm.DB) error {
	fmt.Println("钩子函数AfterDelete触发了")
	return nil
}*/
func (u *UserModel) AfterFind(tx *gorm.DB) (err error) {
	fmt.Println("查询钩子")
	fmt.Println(u)
	return
}
