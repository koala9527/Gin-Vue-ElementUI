package model

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/ini.v1"
)

var db *gorm.DB

func DbConnect() {
	//配置文件路径
	file, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	//读取数据配置
	DBUser := file.Section("database").Key("DbUser").String()
	Password := file.Section("database").Key("DbPassWord").String()
	Host := file.Section("database").Key("DbHost").String()
	Port := file.Section("database").Key("DbPort").MustInt(3306)
	Db := file.Section("database").Key("DbName").String()
	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, Password, Host, Port, Db)
	//创建一个数据库的连接
	db, err = gorm.Open("mysql", connArgs)
	// defer db.Close()
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)                               //去掉数据库自动迁移数据名加的s
	db.AutoMigrate(&UrlList{}, &UrlType{}, &AdminUser{}) //自动迁移
}

func DBClose() {
	if err := db.Close(); err != nil {
		log.Println("数据库关闭失败：" + err.Error())
	}
}

type UrlType struct {
	ID       uint `gorm:"primary_key"`
	Name     string
	UrlLists []UrlList `gorm:"FOREIGNKEY:TypeID;ASSOCIATION_FOREIGNKEY:ID"`
}

type AdminUser struct {
	ID           uint `gorm:"primary_key"`
	Name         string
	Avatar       string
	Password     string
	Token        string
	ExpirationAt time.Time `json:"created_at"`
}

type UrlList struct {
	ID     uint `gorm:"primary_key"`
	TypeID int  // 默认外键， 用户Id
	Name   string
	URL    string
}

func GetList() (data interface{}, err error) {
	var list []UrlType
	db.Debug().Preload("UrlLists").Find(&list)
	return list, nil
}

//获取所有的type
func GetAllTypeLIst() (data interface{}, err error) {
	var list []UrlType
	db.Debug().Find(&list)
	return list, nil
}

//删除一个分类
func DelType(id int) error {
	return db.Delete(&UrlType{}, id).Error
}

//增加一个分类
func AddType(name string) (id int, err error) {
	TypeData := UrlType{
		Name: name,
	}
	err = db.Debug().Create(&TypeData).Error
	return int(TypeData.ID), err
}

//修改一个分类
func EditType(id int, name string) error {
	return db.Debug().Model(&UrlType{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Name": name,
	}).Error
}

func Login(name string, password string) (list []AdminUser, err error) {
	var user []AdminUser
	db.Debug().Where("name = ? and password = ?", name, password).First(&user)

	return user, nil
}

func LoginCreateToken(name string, password string, token string, expiration_at string) (err error) {
	return db.Debug().Table("admin_user").Where("name = ? and password = ?", name, password).Updates(map[string]interface{}{"token": token, "expiration_at": expiration_at}).Error
}

func TokenInfo(token string) (list []AdminUser, err error) {
	var user []AdminUser
	db.Debug().Where("token = ? ", token).First(&user)
	return user, nil
}

func Info(token string) (data interface{}, err error) {
	type Result struct {
		Name   string
		Avatar string
	}
	var result Result
	db.Debug().Table("admin_user").Select("name, avatar").Where("token = ? ", token).Scan(&result)
	return result, nil
}

func Logout(token string) (err error) {
	return db.Debug().Table("admin_user").Where("token = ? ", token).Updates(map[string]interface{}{"token": ""}).Error
}
func UpdateTokenTime(token string, expiration_at string) (err error) {
	return db.Debug().Table("admin_user").Where("token = ? ", token).Updates(map[string]interface{}{"expiration_at": expiration_at}).Error
}

func TokenFind(token string) (status bool) {
	var user []AdminUser
	rowsAffects := db.Debug().Where("token = ? ", token).First(&user).RowsAffected
	if rowsAffects == 0 {
		return false
	} else {
		return true
	}
}

func GetUrlList(id int) (data interface{}, err error) {
	var list []UrlList

	if id != 0 {
		res := db.Debug().Where("type_id = ?", id).Find(&list)
		return res, nil
	} else {
		res := db.Debug().Find(&list)
		return res, nil
	}

}

//删除一个详情
func DelUrl(id int) error {
	return db.Delete(&UrlList{}, id).Error
}

//增加一个详情
func AddUrl(name string, url string, type_id int) (id int, err error) {
	ListData := UrlList{
		Name:   name,
		TypeID: type_id,
		URL:    url,
	}
	err = db.Debug().Create(&ListData).Error
	return int(ListData.ID), err
}

//修改一个详情
func EditUrl(id int, type_id int, name string, url string) error {
	return db.Debug().Model(&UrlList{}).Where("id = ?", id).Updates(map[string]interface{}{
		"TypeID": type_id,
		"Name":   name,
		"Url":    url,
	}).Error
}
