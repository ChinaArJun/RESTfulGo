package user

//CreateRequest 和 CreateResponse，
//并将这些结构体统一放在一个 Go 文件夹中，以方便后期维护和修改

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}
