package route

import (
	"go-wxbot/controller"

	"github.com/gin-gonic/gin"
)

// 初始化消息相关路由
func initMessageRoute(app *gin.Engine) {
	group := app.Group("/message")

	// 向指定好友发送图片消息
	group.POST("/file", controller.SendFileHandle)
	// 向指定好友发送视频消息
	group.POST("/video", controller.SendVideoHandle)
	// 发送图片消息
	group.POST("/img", controller.SendImgHandle)
	// 发送文本消息
	group.POST("/text", controller.SendTextHandle)

}
