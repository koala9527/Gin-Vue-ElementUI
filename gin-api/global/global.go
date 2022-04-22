package global

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Result 暴露
type Result struct {
	Ctx *gin.Context
}

//ResultCont 返回的结果：
type ResultCont struct {
	Code int         `json:"code"` //提示代码
	Msg  string      `json:"msg"`  //提示信息
	Data interface{} `json:"data"` //数据
}

//NewResult 创建新返回对象
func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

//Success 成功
func (r *Result) Success(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := ResultCont{}
	res.Code = 200
	res.Data = data
	res.Msg = "成功！"
	r.Ctx.JSON(http.StatusOK, res)
}

//Error 出错
func (r *Result) Error(code int, msg string, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := ResultCont{}
	res.Code = code
	res.Msg = msg
	res.Data = data
	r.Ctx.JSON(http.StatusOK, res)
}
