// 数据处理  引入config 中 ini 文件
package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string
	JwtKey string

	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string

	AccessKey string
	SecretKey string
	Bucket    string
	QinServer string
)

func init()  {
	file , err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查路径",err)
	}
	LoadServer(file)
	loadData(file)
	LoadQiniu(file)
}

func LoadServer(file *ini.File)  {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js782js72")

}
func loadData(file *ini.File)  {
	Db = file.Section("server").Key("Db").MustString("mysql")
	DbHost = file.Section("server").Key("DbHost").MustString("localhost")
	DbPort = file.Section("server").Key("DbPort").MustString("3306")
	DbUser = file.Section("server").Key("DbUser").MustString("root")
	DbPassWord = file.Section("server").Key("DbPassWord").MustString("root")
	DbName = file.Section("server").Key("DbName").MustString("ginblog")

}
func LoadQiniu(file *ini.File)  {
	AccessKey  =file.Section("qiniu").Key("AccessKey").String()
	SecretKey =file.Section("qiniu").Key("SecretKey").String()
	Bucket =file.Section("qiniu").Key("Bucket").String()
	QinServer =file.Section("qiniu").Key("QinServer").String()

}