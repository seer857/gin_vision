package model

import (
	"fmt"
	"ginblog/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

//连接配置数据库
var db *gorm.DB
var err error
func InitDb()  {
 	db,err = gorm.Open(utils.Db,fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
 	utils.DbUser,
 	utils.DbPassWord,
 	utils.DbHost,
 	utils.DbPort,
 	utils.DbName,
 	))
 	if err != nil{
 		fmt.Printf("连接数据库失败，请检查参数",err)
	}
	// 禁用默认表面的复数形式
	db.SingularTable(true)
	//创建的ORM模型 自动迁移
	db.AutoMigrate(&User{},&Article{},&Category{},&Sale{},&SaleThisYear{},&CarSales{})
	// SetMaxIdleConns 设置连接池中的最大闲置连接数
	db.DB().SetMaxIdleConns(10)
 	// SetMaxOpenConns 设置数据库组大连接数量
	db.DB().SetMaxOpenConns(100)
 	// SetConnMaxLifetime 设置连接最大可复用时间
 	db.DB().SetConnMaxLifetime(10 * time.Second)

 	//db.Close()
}