package route

import (
	"go-wxbot/controller"

	"github.com/gin-gonic/gin"
)

// 初始化消息相关路由
func initMessageRoute(app *gin.Engine) {
	group := app.Group("/message")

	// 群发文本消息
	group.POST("/batch", controller.SendTextBatchHandle)

	// 向指定群组发送视频消息
	group.POST("/group/video", controller.SendVideoToGroupHandle)

	// 向指定好友发送视频消息
	group.POST("/user/video", controller.SendVideoToFriendHandle)

	// 向指定群组发送图片消息
	group.POST("/group/file", controller.SendFileToGroupHandle)

	// 向指定好友发送图片消息
	group.POST("/user/file", controller.SendFileToFriendHandle)

	// 向指定群组发送图片消息
	group.POST("/group/img", controller.SendImgToGroupHandle)

	// 向指定好友发送图片消息
	group.POST("/user/img", controller.SendImgToFriendHandle)

	// 向指定好友发送文本消息
	group.POST("/user", controller.SendTextToFriendHandle)

	// 向指定群组发送消息
	group.POST("/group", controller.SendTextToGroupHandle)
}
