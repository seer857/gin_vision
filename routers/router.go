package routers

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/routers/cors"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter()  {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.Use(cors.CORS())
	auth := r.Group("api/v1")
		auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		auth.POST("user/add",v1.AddUser)
		auth.PUT("user/:id",v1.EditUser)
		auth.DELETE("user/:id",v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add",v1.AddCategory)
		auth.PUT("category/:id",v1.EditCate)
		auth.DELETE("category/:id",v1.DeleteCate)
		//文章模块的路由接口
		auth.POST("article/add",v1.AddArticle)
		auth.PUT("article/:id",v1.EditArt)
		auth.DELETE("article/:id",v1.DeleteArt)
		// 上传文件
		auth.POST("upload",v1.Upload)
	}
	router :=r.Group("api/v1")
	{
		// 用户模块的路由接口
		router.GET("users",v1.GetUsers)
		// 分类模块的路由接口
		router.GET("category",v1.GetCate)
		//文章模块的路由接口
		router.GET("article",v1.GetArt)
		router.GET("article/list/:id",v1.GetCateArt)
		router.GET("article/info/:id",v1.GetArtInfo)
		//查询销量路由
		router.GET("sales",v1.GetSalesAll)
		router.GET("sale21",v1.GetSalesThisYear)
		// 车型销量路由
		router.GET("carsales",v1.GetCarSales)
		router.POST("login",v1.Login)
	}

	r.Run(utils.HttpPort)
}