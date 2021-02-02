package model

type CarSales struct {
	id int `gorm:type:int; json:"id"`
	VehicleSeriesDesc string `gorm:"type:varchar(20) ;json:"vehicle_series_desc"`
	ActualQuantityDelivered string `gorm:"type:varchar(20)"; json:"actual_quantity_delivered" `
	MulPrice string `gorm:"type:varchar(20)" ;json:"mul_price" `
}

// 查询 车型信息表
func GetCarSalesAll(pageSize int,pageNum int)[]CarSales  {
	var carSales []CarSales
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&carSales).Error
	if err != nil{
		return nil
	}
	return carSales
}