package request

// 定义接收数据的结构体
type DelTypeRequest struct {
	Id int `form:"id" json:"id" binding:"required"`
}

type AddTypeRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type EditTypeRequest struct {
	Id   int    `form:"id" json:"id" binding:"required"`
	Name string `form:"name" json:"name" binding:"required"`
}

type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type GetUrlListRequest struct {
	Id int `form:"id" json:"id"`
}

type DelUrlRequest struct {
	Id int `form:"id" json:"id" binding:"required"`
}

type AddUrlRequest struct {
	Name   string `form:"name" json:"name" binding:"required"`
	TypeID int    `form:"type_id" json:"type_id" binding:"required"`
	Url    string `form:"url" json:"url" binding:"required"`
}

type EditUrlRequest struct {
	Id     int    `form:"id" json:"id" binding:"required"`
	Name   string `form:"name" json:"name" binding:"required"`
	TypeID int    `form:"type_id" json:"type_id" binding:"required"`
	Url    string `form:"url" json:"url" binding:"required"`
}
