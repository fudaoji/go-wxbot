package route

import (
	"go-wxbot/controller"

	"github.com/gin-gonic/gin"
)

// initUserRoute 初始化登录路由信息
func initUserRoute(app *gin.Engine) {
	group := app.Group("/user")
	// 移除多个群成员
	group.POST("/group/removemembers", controller.RemoveMembersFromGroupHandle)
	// 获取群成员列表
	group.POST("/group/members", controller.GetGroupMembersHandle)
	// 邀请好友入多群
	group.POST("/addfriendintogroups", controller.AddFriendIntoGroupsHandle)

	// 邀请多个好友入群
	group.POST("/addfriendsintogroup", controller.AddFriendsIntoGroupHandle)
	// 获取登录的用户信息
	group.GET("/info", controller.GetCurrentUserInfoHandle)
	// 获取好友列表
	group.GET("/friends", controller.GetFriendsListHandle)
	// 获取群组列表
	group.GET("/groups", controller.GetGroupsListHandle)
	// 修改好友备注
	group.POST("/setremarkname", controller.SetRemarkNameHandle)
}
