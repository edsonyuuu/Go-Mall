package routers

import (
	"Go-Mall/AppV1/api"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var store = cookie.NewStore([]byte("HyvCD89g3VDJ9646BFGEh37GFJ"))

func (router RouterGroup) UserRouter() {
	app := api.ApiGroup.UserApi
	router.Use(sessions.Sessions("sessionid", store))
	router.POST("user_register", app.Register)
}
