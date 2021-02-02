package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCarSales(c *gin.Context)  {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0{
		pageNum = -1
	}
	data := model.GetCarSalesAll(pageSize,pageNum)
	code = errmsg.SUCCSE
	c.JSON(http.StatusOK,gin.H{
		"status": code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}