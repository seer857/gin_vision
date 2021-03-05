package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context)  {
	file,fileHeader,_ := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	url,code :=model.UploadFile(file,fileSize)
	c.JSON(http.StatusOK,gin.H{
		"stauts":code,
		"messages":errmsg.GetErrMsg(code),
		"url":url,
	})
}
