package route

import (
	"go-wxbot/controller"

	"github.com/gin-gonic/gin"
)

// initUserRoute 初始化登录路由信息
func initUserRoute(app *gin.Engine) {
	group := app.Group("/user")

	// 获取登录的用户信息
	group.GET("/info", controller.GetCurrentUserInfoHandle)
	// 获取好友列表
	group.GET("/friends", controller.GetFriendsListHandle)
	// 获取群组列表
	group.GET("/groups", controller.GetGroupsListHandle)
	// 修改好友备注
	group.POST("/setfriendremarkname", controller.SetFriendRemarkNameHandle)
}
