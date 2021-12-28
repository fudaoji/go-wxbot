package route

import (
	"go-wxbot/controller"

	"github.com/gin-gonic/gin"
)

// initLoginRoute 初始化登录路由信息
func initLoginRoute(app *gin.Engine) {
	// 获取登录二维码
	app.GET("/getlogincode", controller.GetLoginUrlHandle)
	// 检查登录状态
	app.POST("/checklogin", controller.LoginHandle)
}
