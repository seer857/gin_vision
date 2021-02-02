package model

import "github.com/jinzhu/gorm"

type Sale struct {
	//gorm.Model  // 默认带3个参数  创建时间 更新时间 删除时间
	id int `gorm:type:int; json:"id"`
	Delivery_ym string `gorm:"type:varchar(20);" json:"delivery_ym"`
	Actual_quantity_delivered_bq string `gorm:"type:varchar(20);" json:"actual_quantity_delivered_bq"`
	Mul_price string `gorm:"type:"DECIMAL"; json:"mul_price"`
}
type SaleThisYear struct {
	//gorm.Model  // 默认带3个参数  创建时间 更新时间 删除时间
	id int `gorm:type:int; json:"id"`
	Delivery_ym string `gorm:"type:varchar(20);" json:"delivery_ym"`
	Actual_quantity_delivered_bq string `gorm:"type:varchar(20);" json:"actual_quantity_delivered_bq"`
	Mul_price string `json:"mul_price"`
}

func GetSalesAll(pageSize int,pageNum int)[]Sale  {
	var sales []Sale
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Order("id").Find(&sales).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return sales
}
func GetSalesThisYear(pageSize int,pageNum int)[]SaleThisYear  {
	var saleThisYear []SaleThisYear
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Order("id").Find(&saleThisYear).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return saleThisYear
}