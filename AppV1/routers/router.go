package routers

import (
	"Go-Mall/AppV1/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	//设置模式
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	//路由群
	apiRouterGroup := router.Group("api")
	routerGroupApp := RouterGroup{apiRouterGroup}
	//用户操作
	routerGroupApp.UserRouter()
	//用户对商品操作

	//用户需要鉴权进行操作

	//管理员操作商品

	//收藏夹

	//购物车

	//订单操作

	//收货地址

	//支付功能

	//显示余额

	//秒杀专场

	return router
}
