package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func AddArticle(c *gin.Context)  {
	// todo 添加用户
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	code = model.CreateArt(&data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
//  todo 查询分类下的所有文章

// 查询文章列表
func GetArt(c *gin.Context)  {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0{
		pageNum = -1
	}
	data := model.GetCate(pageSize,pageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
// 编辑文章
func EditArt(c *gin.Context)  {
	var data model.Article
	id,_ := strconv.Atoi(c.Query("id"))
	c.ShouldBindJSON(&data)
	code = model.EditArt(id,&data)

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
// 删除文章
func DeleteArt(c *gin.Context)  {
	// id,_ 接受传回来的参数
	id,_ := strconv.Atoi(c.Query("id"))
	code = model.DeleteArt(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),

	})

}