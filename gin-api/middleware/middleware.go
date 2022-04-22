package middleware

import (
	"fmt"
	"main/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,x-token,token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
func Token() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("X-Token")
		token_exsits, err := model.TokenInfo(token)
		if err != nil {
			resospnseWithError(401, "非法请求", context)
			return
		}
		fmt.Println(token_exsits)
		if len(token_exsits) != 0 {
			//先做时间判断
			target_time, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Time(token_exsits[0].ExpirationAt).Format("2006-01-02 15:04:05"), time.Local) //需要加上time.Local不然会自动加八小时
			if target_time.Unix() <= time.Now().Unix() {
				fmt.Println("过期报错")
				//过期报错
				resospnseWithError(401, "timeout", context)
				return
			}
			//token没过期，更新到期时间
			now := time.Unix(time.Now().Unix()+7200, 0).Format("2006-01-02 15:04:05")
			err = model.UpdateTokenTime(token, now)
			context.Set("name", token_exsits[0].Name)
			context.Set("avatar", token_exsits[0].Avatar)
			context.Set("token", token)
		} else {
			fmt.Println("没了")
			resospnseWithError(401, "已退出", context)
			return
		}

		context.Next()
	}
}

type ResultCont struct {
	Code int         `json:"code"` //提示代码
	Msg  string      `json:"msg"`  //提示信息
	Data interface{} `json:"data"` //数据
}

func resospnseWithError(code int, message string, c *gin.Context) {
	var res ResultCont
	res.Code = code
	res.Msg = message
	c.JSON(200, res) //前端返回也要返回200才能拦截
	c.Abort()
}
