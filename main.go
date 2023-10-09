package main

import (
	core2 "Go-Mall/AppV1/core/core"
	"Go-Mall/AppV1/global"
	"Go-Mall/AppV1/routers"
)

func main() {
	//读取配置文件
	core2.InitConf()
	//初始化日志
	global.Log = core2.InitLogger()
	//连接数据库
	global.DB = core2.InitGorm()

	core2.InitAddrDB()
	defer global.AddrDB.Close()

	//连接redis
	global.Redis = core2.ConnectRedis()
	//连接es
	global.ESClient = core2.EsConnect()

	router := routers.InitRouter()

	addr := global.Config.System.Addr()
	global.Log.Infof("Go-Mall 运行在：%s", addr)
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
