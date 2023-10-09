package model

type User struct {
	Base
	UserName string `json:"user_name" form:"user_name" comment:"用户名" gorm:"unique"`
	NickName string `json:"nick_name" form:"nick_name" comment:"昵称"`
	Password string `json:"password" form:"password" comment:"密码"`
	Email    string `json:"email" form:"email" comment:"邮箱"`
}
