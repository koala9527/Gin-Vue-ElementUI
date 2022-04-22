package controller

import (
	"main/global"
	"main/request"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

//
func Api(c *gin.Context) {

	result := global.NewResult(c)
	data, err := service.Api()
	if err != nil {
		result.Error(5201, err.Error(), data)
		return
	}
	result.Success(data)
}

func GetTypeList(c *gin.Context) {
	result := global.NewResult(c)
	data, err := service.GetTypeList()
	if err != nil {
		result.Error(5201, err.Error(), data)
		return
	}
	result.Success(data)
}

func DelType(c *gin.Context) {
	var json request.DelTypeRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := global.NewResult(c)
	err := service.DelType(json)
	if err != nil {
		result.Error(5201, err.Error(), "删除失败")
		return
	}
	result.Success("删除成功")
}

func AddType(c *gin.Context) {
	var json request.AddTypeRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := global.NewResult(c)
	id, err := service.AddType(json)
	if err != nil {
		result.Error(5201, err.Error(), "添加成功")
		return
	}
	result.Success(id)
}

func EditType(c *gin.Context) {
	var json request.EditTypeRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := global.NewResult(c)
	err := service.EditType(json)
	if err != nil {
		result.Error(5201, err.Error(), "修改失败")
		return
	}
	result.Success("修改成功")
}

func Login(c *gin.Context) {
	var json request.LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := global.NewResult(c)
	data, err := service.Login(json)
	if err != nil {
		result.Error(5201, err.Error(), "登录失败")
		return
	}
	result.Success(data)
}

func Info(c *gin.Context) {
	token := c.MustGet("token").(string)
	result := global.NewResult(c)
	data, err := service.Info(token)
	if err != nil {
		result.Error(5201, err.Error(), "登录失败")
		return
	}
	result.Success(data)
}

func Logout(c *gin.Context) {
	token := c.MustGet("token").(string)
	result := global.NewResult(c)
	err := service.Logout(token)
	if err != nil {
		result.Error(5201, err.Error(), "退出登录失败")
		return
	}
	result.Success("退出登录成功")
}

func GetUrlList(c *gin.Context) {
	var json request.GetUrlListRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := global.NewResult(c)
	list, err := service.GetUrlList(json)
	if err != nil {
		result.Error(5201, err.Error(), "域名获取失败")
		return
	}
	result.Success(list)
}

func DelUrl(c *gin.Context) {
	var json request.DelUrlRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := global.NewResult(c)
	err := service.DelUrl(json)
	if err != nil {
		result.Error(5201, err.Error(), "删除失败")
		return
	}
	result.Success("删除成功")
}

func AddUrl(c *gin.Context) {
	var json request.AddUrlRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := global.NewResult(c)
	id, err := service.AddUrl(json)
	if err != nil {
		result.Error(5201, err.Error(), "域名获取失败")
		return
	}
	result.Success(id)
}

func EditUrl(c *gin.Context) {
	var json request.EditUrlRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result := global.NewResult(c)
	err := service.EditUrl(json)
	if err != nil {
		result.Error(5201, err.Error(), "域名修改失败")
		return
	}
	result.Success("域名修改成功")
}
