package routers

import (
	v1 "ginblog/api/v1"
	"ginblog/routers/cors"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.Use(cors.CORS())
	router := r.Group("api/v1")
	{
		// 用户模块的路由接口
		router.POST("user/add",v1.AddUser)
		router.GET("users",v1.GetUsers)
		router.PUT("user/:id",v1.EditUser)
		router.DELETE("user/:id",v1.DeleteUser)
		// 分类模块的路由接口

		//文章模块的路由接口

		//查询销量路由
		router.GET("sales",v1.GetSalesAll)
		router.GET("sale21",v1.GetSalesThisYear)
		// 车型销量路由
		router.GET("carsales",v1.GetCarSales)
		// 跨域
		
	}
	r.Run(utils.HttpPort)
}