package user_api

import "github.com/gin-gonic/gin"

type UserRegister struct {
	NickName string `json:"nick_name" form:"nick_name"`
	UserName string `json:"user_name" form:"user_name"`
	Password string `json:"password" form:"password"`
}

func (UserApi) Register(c *gin.Context) {

}
