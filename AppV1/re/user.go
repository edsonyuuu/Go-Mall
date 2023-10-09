package re

type UserRegisterReq struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}

type UserLoginReq struct {
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}
