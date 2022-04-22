package router

import (
	"log"
	"main/controller"
	"main/global"
	"main/middleware"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

//Router 路由方法
func Router() *gin.Engine {
	router := gin.Default()
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(Recover)
	router.Use(middleware.Cors())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, 1)
		return
	})

	router.GET("/api", controller.Api)
	admin := router.Group("/admin")
	admin.Use(middleware.Token())
	{
		// 路径映射
		// api:=controller.NewDyController()
		admin.GET("/getTypeList", controller.GetTypeList)
		admin.POST("/DelType", controller.DelType)
		admin.POST("/AddType", controller.AddType)
		admin.POST("/EditType", controller.EditType)
		admin.POST("/getUrlList", controller.GetUrlList)
		admin.POST("/DelUrl", controller.DelUrl)
		admin.POST("/AddUrl", controller.AddUrl)
		admin.POST("/EditUrl", controller.EditUrl)
		admin.POST("/user/logout", controller.Logout)
		admin.GET("/user/info", controller.Info)
	}

	adminuser := router.Group("/admin/user")
	{
		adminuser.POST("/login", controller.Login)
	}

	return router
}

//HandleNotFound 404
func HandleNotFound(c *gin.Context) {
	global.NewResult(c).Error(404, "资源未找到", nil)
	return
}

//Recover 500
func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			global.NewResult(c).Error(500, "服务器内部错误", nil)
		}
	}()
	//继续后续接口调用
	c.Next()
}
