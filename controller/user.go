package controller

import (
	"go-wxbot/core"
	"go-wxbot/global"
	"go-wxbot/logger"
	"go-wxbot/protocol"

	"github.com/fudaoji/go-utils"

	"github.com/eatmoreapple/openwechat"
	"github.com/gin-gonic/gin"
)

// 返回用户信息包装类
type responseUserInfo struct {
	Uin         int64              `json:"uin"`          // 用户唯一ID
	Sex         int                `json:"sex"`          // 性别
	Province    string             `json:"province"`     // 省
	City        string             `json:"city"`         // 市
	Alias       string             `json:"alias"`        // 别名
	DisplayName string             `json:"display_name"` // 显示名称
	NickName    string             `json:"nick_name"`    // 昵称
	RemarkName  string             `json:"remark_name"`  // 备注
	HeadImgUrl  string             `json:"head_img_url"` // 头像
	UserName    string             `json:"user_name"`    // 当前登录中用户的唯一标识
	Members     []*openwechat.User `json:"members"`      // 群成员(群独有)
}

// 返回的好友列表的实体
type friendsResponse struct {
	Total   int                `json:"total"`
	Friends []responseUserInfo `json:"friends"`
}

// 返回的群聊列表的实体
type groupsResponse struct {
	Total  int                `json:"total"`
	Groups []responseUserInfo `json:"groups"`
}

// 修改备注名请求体
type setRemarkNameRes struct {
	// 用户名
	To string `form:"to" json:"to"`
	// 正文
	RemarkName string `form:"remark_name" json:"remark_name"`
}

// 邀请好友入群请求体
type addFriendsIntoGroupRes struct {
	// 群username
	Group string `form:"group" json:"group"`
	// 好友username数组
	Friends []string `form:"friends" json:"friends"`
}

// 邀请好友入多群请求体
type addFriendIntoGroupsRes struct {
	// 好友username
	Friend string `form:"friend" json:"friend"`
	// 群组username数值
	Groups []string `form:"groups" json:"groups"`
}

// 获取群成员请求体
type getGroupMembersRes struct {
	// 群username
	Group string `form:"group" json:"group"`
}

// 获取群成员请求体
type getGroupMembersResp struct {
	Total   int                `json:"total"`
	Members []responseUserInfo `json:"members"` // 群成员
}

// 移出群请求体
type removeMembersFromGroupRes struct {
	// 群username
	Group string `form:"group" json:"group"`
	// 好友username数组
	Members []string `form:"members" json:"members"`
}

// RemoveMembersFromGroupHandle 将好友移除群
func RemoveMembersFromGroupHandle(ctx *gin.Context) {
	// 取出请求参数
	var res removeMembersFromGroupRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	group, self := FindGroup(bot, res.Group, ctx)
	if group == nil {
		return
	}

	var members = make(openwechat.Members, 0)
	allMembers, _ := group.Members()
	for _, item := range allMembers {
		if exists, _ := utils.ContainsStr(res.Members, item.UserName); exists {
			members = append(members, item)
		}
	}
	if len(members) > 0 {
		if err := self.RemoveMemberFromGroup(group, members); err != nil {
			core.FailWithMessage("移除群成员出错："+err.Error(), ctx)
			return
		}
	}
	core.Ok(ctx)
}

// GetGroupMembersHandle 获取群成员
func GetGroupMembersHandle(ctx *gin.Context) {
	// 取出请求参数
	var res getGroupMembersRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}
	bot := GetCurBot(ctx)
	group, _ := FindGroup(bot, res.Group, ctx)
	if group == nil {
		return
	}
	// 获取好友列表
	members, err := group.Members()
	if err != nil {
		core.FailWithMessage("获取好友列表失败", ctx)
		return
	}

	// 循环处理数据
	var memberList []responseUserInfo
	for _, friend := range members {
		memberList = append(memberList, responseUserInfo{
			Uin:         friend.Uin,
			Sex:         friend.Sex,
			Province:    friend.Province,
			City:        friend.City,
			Alias:       friend.Alias,
			DisplayName: friend.DisplayName,
			NickName:    friend.NickName,
			RemarkName:  friend.RemarkName,
			HeadImgUrl:  friend.HeadImgUrl,
			UserName:    friend.UserName,
		})
	}

	// 返回给前端
	core.OkWithData(getGroupMembersResp{Total: len(memberList), Members: memberList}, ctx)
}

// AddFriendsIntoGroupHandle 邀请1个好友入多个群
func AddFriendIntoGroupsHandle(ctx *gin.Context) {
	// 取出请求参数
	var res addFriendIntoGroupsRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	friend, self := FindFriend(bot, res.Friend, ctx)
	if friend == nil {
		return
	}

	var groups = make(openwechat.Groups, 0)
	for _, item := range res.Groups {
		group, _ := FindGroup(bot, item, ctx)
		if friend != nil {
			groups = append(groups, group)
		}
	}

	self.AddFriendIntoManyGroups(friend, groups...)
	core.Ok(ctx)
}

// AddFriendsIntoGroupHandle 邀请好友入群
func AddFriendsIntoGroupHandle(ctx *gin.Context) {
	// 取出请求参数
	var res addFriendsIntoGroupRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	group, self := FindGroup(bot, res.Group, ctx)
	if group == nil {
		return
	}

	var friends = make(openwechat.Friends, 0)
	for _, item := range res.Friends {
		friend, _ := FindFriend(bot, item, ctx)
		if friend != nil {
			friends = append(friends, friend)
		}
	}

	self.AddFriendsIntoGroup(group, friends...)
	core.Ok(ctx)
}

// SetRemarkNameHandle 修改指定用户的备注
func SetRemarkNameHandle(ctx *gin.Context) {
	// 取出请求参数
	var res setRemarkNameRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}
	bot := GetCurBot(ctx)
	// 获取登录用户
	self, _ := bot.GetCurrentUser()
	// 查找指定的好友
	friends, _ := self.Friends(true)
	// 查询指定好友
	friendSearchResult := friends.SearchByUserName(1, res.To)
	if friendSearchResult.Count() < 1 {
		core.FailWithMessage("指定好友不存在", ctx)
		return
	}
	// 取出好友
	friend := friendSearchResult.First()
	// 设置备注
	if err := self.SetRemarkNameToFriend(friend, res.RemarkName); err != nil {
		core.FailWithMessage("设置备注失败："+err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

// GetCurrentUserInfoHandle 获取当前登录用户
func GetCurrentUserInfoHandle(ctx *gin.Context) {
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")
	bot := global.GetBot(appKey)
	// 获取登录用户信息
	user, err := bot.GetCurrentUser()
	if err != nil {
		core.FailWithMessage("获取登录用户信息失败", ctx)
		return
	}

	logger.Log.Infof("登录用户：%v", user.NickName)
	//下载头像
	userDataVO := responseUserInfo{
		Uin:         user.Uin,
		Sex:         user.Sex,
		Province:    user.Province,
		City:        user.City,
		Alias:       user.Alias,
		DisplayName: user.DisplayName,
		NickName:    user.NickName,
		RemarkName:  user.RemarkName,
		HeadImgUrl:  user.HeadImgUrl,
		UserName:    user.UserName,
	}
	core.OkWithData(userDataVO, ctx)
}

// GetFriendsListHandle 获取好友列表
func GetFriendsListHandle(ctx *gin.Context) {
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")

	bot := global.GetBot(appKey)
	// 获取好友列表
	user, _ := bot.GetCurrentUser()
	friends, err := user.Friends(true)
	if err != nil {
		core.FailWithMessage("获取好友列表失败", ctx)
		return
	}

	// 循环处理数据
	var friendList []responseUserInfo
	for _, friend := range friends {
		friendList = append(friendList, responseUserInfo{
			Uin:         friend.Uin,
			Sex:         friend.Sex,
			Province:    friend.Province,
			City:        friend.City,
			Alias:       friend.Alias,
			DisplayName: friend.DisplayName,
			NickName:    friend.NickName,
			RemarkName:  friend.RemarkName,
			HeadImgUrl:  friend.HeadImgUrl,
			UserName:    friend.UserName,
		})
	}

	// 返回给前端
	core.OkWithData(friendsResponse{Total: friends.Count(), Friends: friendList}, ctx)
}

// GetGroupsListHandle 获取群聊列表
func GetGroupsListHandle(ctx *gin.Context) {
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")

	bot := global.GetBot(appKey)
	// 获取好友列表
	user, _ := bot.GetCurrentUser()

	groups, err := user.Groups(true)
	if err != nil {
		core.FailWithMessage("获取群聊列表失败", ctx)
		return
	}

	logger.Log.Infof("%v", groups)
	// 循环处理数据
	var groupList []responseUserInfo
	for _, group := range groups {
		// 取出群成员
		//members, _ := group.Members()
		groupList = append(groupList, responseUserInfo{
			Uin:         group.Uin,
			Sex:         group.Sex,
			Province:    group.Province,
			City:        group.City,
			Alias:       group.Alias,
			DisplayName: group.DisplayName,
			NickName:    group.NickName,
			RemarkName:  group.RemarkName,
			HeadImgUrl:  group.HeadImgUrl,
			UserName:    group.UserName,
			//Members:     members,
		})
	}

	// 返回给前端
	core.OkWithData(groupsResponse{Total: groups.Count(), Groups: groupList}, ctx)
}

//GetBot 获取当前bot
func GetCurBot(ctx *gin.Context) *protocol.WechatBot {
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")
	return global.GetBot(appKey)
}
