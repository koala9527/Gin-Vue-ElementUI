package main

import (
	"fmt"
	"main/model"
	"main/router"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/ini.v1"
)

var HttpPort string
var AppMode string

type config struct {
	App string
}

func init() {
	//配置文件路径
	file, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	//读取项目运行
	HttpPort = file.Section("server").Key("HttpPort").String()
	AppMode = file.Section("server").Key("AppMode").String()
}

func main() {
	model.DbConnect()
	defer model.DBClose()
	r := router.Router()
	gin.SetMode(AppMode)
	r.Run(HttpPort)
}
