package model

import (
	"ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model  // 默认带3个参数  创建时间 更新时间 删除时间
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role int `gorm:"type:int" json:"role"`
	//Avatar String
}
// 查询用户是否存在
func CheckUser(name string)(code int)  {
	var users User
	db.Select("id").Where("username == ?",name).First(&users)
	if users.ID > 0{
		return errmsg.ERROR_USERNAME_USED // 1001
	}
	return errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *User)int  {
	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// 查询用户列表
func GetUsers(pageSize int,pageNum int)[]User  {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return users
}

// 编辑用户

// 删除用户