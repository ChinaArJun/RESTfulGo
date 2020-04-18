package user

import (
	"RESTfulGo/model"
)

//CreateRequest 和 CreateResponse，
//并将这些结构体统一放在一个 Go 文件夹中，以方便后期维护和修改

type ListRequest struct {
	Pagesize int `json:"pagesize"`
	PageNum int `json:"pagenum"`
}

type ListResponse struct {
	TotalCount int64 `json:"total_count"`
	UserList []*model.UserInfo `json:"user_list"`
}

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type GetUserResponse struct {
	Id string `json:"id"`
}