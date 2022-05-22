package models

/*
* 用户主表
 */
import (
	"github.com/pkg/errors"
	"warp-demo/basic/db"
)

type Users struct {
	Id    int64  `gorm:"primary_key" json:"id"`
	Name  string `gorm:"column:name;type:varchar(255)" json:"name" description:"昵称"`
	Phone string `gorm:"column:phone;type:varchar(255)" json:"phone" description:"手机号"`
}

func NewUsers() *Users {
	return &Users{}
}
func (u Users) TableName() string {
	return "users"
}

// GetUserById /*根据用户id查询*/
func (u *Users) GetUserById(userid int64) (v *Users, err error) {
	o := db.GetDB()
	v = &Users{}
	if err := o.Where(Users{Id: userid}).First(&v).Error; err != nil {
		return v, errors.Wrap(err, "users.go GetUserById get user failed\n")
	}
	return
}
