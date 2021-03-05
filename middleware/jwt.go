package middleware

import (
	"ginblog/utils"
	"ginblog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// 中间件
// JwtKey 用来生成Token命令 byte切片
var JwtKey = []byte(utils.JwtKey)

type MyClamis struct {
	Username string  `json:"username"`
	jwt.StandardClaims
}
// 1、生成token
func SetToken(username string)(string,int){
	expireTime := time.Now().Add(10*time.Hour)
	SetClaims := MyClamis{
		username,
		 jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "ginblog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,SetClaims)
	token,err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "",errmsg.ERROR
	}
	return token,errmsg.SUCCSE
}
// 2、验证token
func CheckToken(token string)(*MyClamis,int)  {
	setToken,_ := jwt.ParseWithClaims(token,&MyClamis{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey,nil
	})
	if key,_ :=setToken.Claims.(*MyClamis);setToken.Valid{
		return key,errmsg.SUCCSE
	}else {
		return nil,errmsg.ERROR
	}

}

// jwt中间件
func JwtToken()gin.HandlerFunc  {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCSE
		if tokenHerder == ""{
			code = errmsg.ERROR_TOKEN_EXIST
			//c.Abort()
			//return
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"massage":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHerder," ",2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer"{
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"massage":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key,tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"massage":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix()>key.ExpiresAt{
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK,gin.H{
				"code":code,
				"massage":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username",key.Username)
		c.Next()
	}
}
