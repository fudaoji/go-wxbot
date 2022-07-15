package main

import (
	"go-wxbot/core"
	"go-wxbot/handler"

	"github.com/gin-gonic/gin"

	. "go-wxbot/db"
	"go-wxbot/global"
	"go-wxbot/logger"
	"go-wxbot/middleware"

	. "go-wxbot/install"
	"go-wxbot/route"
)

// 程序启动入口
func main() {
	//读取配置
	core.InitConfig()
	if core.GetIntVal("app_debug", 0) == 0 {
		gin.SetMode(gin.ReleaseMode)
	}

	// 初始化日志
	logger.InitLogger()
	// 初始化Gin
	app := gin.Default()

	// 定义全局异常处理
	app.NoRoute(core.NotFoundErrorHandler())
	// AppKey预检
	app.Use(middleware.CheckAppKeyExistMiddleware(), middleware.CheckAppKeyIsLoggedInMiddleware())
	// 初始化路由
	route.InitRoute(app)

	// 初始化WechatBotMap
	global.InitWechatBotsMap()

	// 初始化MysqlDB
	InitMysqlConnHandle()
	// 安装数据表
	InstallHandle()

	// 初始化Redis连接
	InitRedisConnHandle()

	// 初始化Redis里登录的数据
	handler.InitBotWithStart()

	// 定时更新 Bot 的热登录数据
	global.UpdateHotLoginData()

	// 保活
	global.KeepAliveHandle()

	// 监听端口
	_ = app.Run(":8889")
}
