package service

import (
	"errors"
	"fmt"
	"main/model"
	"main/request"
	"main/util"
	"time"
)

func Api() (data interface{}, err error) {
	list, nil := model.GetList()
	return list, nil
}

func GetTypeList() (data interface{}, err error) {
	list, nil := model.GetAllTypeLIst()
	return list, nil
}

func DelType(json request.DelTypeRequest) (err error) {
	nil := model.DelType(json.Id)
	return nil
}

func AddType(json request.AddTypeRequest) (id int, err error) {
	id, nil := model.AddType(json.Name)
	return id, nil
}

func EditType(json request.EditTypeRequest) (err error) {
	nil := model.EditType(json.Id, json.Name)
	return nil
}

func Login(json request.LoginRequest) (data string, err error) {
	//判断有没有这个用户密码
	//没有报错,有了就判断token有没有，没有就创建返回，有就判断时间，时间没过期就返回，过期了就重新生成返回
	var user []model.AdminUser
	json.Password = util.FixMd5(json.Password + "boomxiakalaka")
	user, err = model.Login(json.Username, json.Password)
	if len(user) > 0 {
		//成功
		if user[0].Token == "" {
			fmt.Println("token是空")
			//token为空，创建更新返回
			now := time.Unix(time.Now().Unix()+7200, 0).Format("2006-01-02 15:04:05")
			token := util.GetRandomString(32)
			err = model.LoginCreateToken(json.Username, json.Password, token, now)
			return token, err
		} else {
			//token不为空 ，判断时间，时间不过期直接返回
			target_time, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Time(user[0].ExpirationAt).Format("2006-01-02 15:04:05"), time.Local) //需要加上time.Local不然会自动加八小时
			if target_time.Unix() >= time.Now().Unix() {
				return user[0].Token, nil
			} else {
				//时间过期
				token := util.GetRandomString(32)
				now := time.Unix(time.Now().Unix()+7200, 0).Format("2006-01-02 15:04:05")
				err = model.LoginCreateToken(json.Username, json.Password, token, now)
				return token, err
			}
		}
	}
	return "登陆失败", errors.New("登陆失败")
}

func Info(token string) (data interface{}, err error) {
	list, err := model.Info(token)
	return list, err
}

func Logout(token string) (err error) {
	err = model.Logout(token)
	return err
}

func GetUrlList(json request.GetUrlListRequest) (data interface{}, err error) {

	list, nil := model.GetUrlList(json.Id)
	return list, nil
}

func DelUrl(json request.DelUrlRequest) (err error) {

	err = model.DelUrl(json.Id)
	return nil
}

func AddUrl(json request.AddUrlRequest) (id int, err error) {

	id, nil := model.AddUrl(json.Name, json.Url, json.TypeID)
	return id, nil
}
func EditUrl(json request.EditUrlRequest) (err error) {

	nil := model.EditUrl(json.Id, json.TypeID, json.Name, json.Url)
	return nil
}
