package model

import (
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"golang.org/x/net/context"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var SecreKey = utils.SecretKey

var Bucket = utils.Bucket
var ImgUrl =utils.QinServer

func UploadFile(file multipart.File,fileSize int64)(string,int)  {
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey,SecreKey)
	upToken := putPolicy.UploadToken(mac)

	cfg:=storage.Config{
		Zone: &storage.ZoneHuabei,
		UseCdnDomains: false,
		UseHTTPS: false,
	}
	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(),&ret,upToken,file,fileSize,&putExtra)
	if err != nil{
		return "",errmsg.ERROR
	}
	url := ImgUrl + ret.Key
	return url,errmsg.SUCCSE
}