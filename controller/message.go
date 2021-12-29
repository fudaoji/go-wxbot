package controller

import (
	"fmt"
	"go-wxbot/core"
	"go-wxbot/global"
	"go-wxbot/protocol"
	"math/rand"
	"os"
	"path"
	"time"

	"github.com/fudaoji/go-utils"

	"github.com/eatmoreapple/openwechat"
	"github.com/gin-gonic/gin"
)

// 发送消息请求体
type sendMsgRes struct {
	// username
	To string `form:"to" json:"to"`
	// 消息类型
	Type string `form:"type" json:"type"`
	// 正文
	Content string `form:"content" json:"content"`
}

// 发送文件请求体
type sendFileRes struct {
	// username
	To string `form:"to" json:"to"`
	// 消息类型
	Type string `form:"type" json:"type"`
	// 正文
	Content string `form:"content" json:"content"`
	// 文件名称
	Filename string `form:"filename" json:"filename"`
}

// 批量发送文本消息请求体
type sendTextBatchRes struct {
	// 好友数组
	Friends []string `form:"friends" json:"friends"`
	// 群聊数组
	Groups []string `form:"groups" json:"groups"`
	// 正文
	Content string `form:"content" json:"content"`
}

// 批量发送文本消息请求体
type sendFileBatchRes struct {
	// 好友数组
	Friends []string `form:"friends" json:"friends"`
	// 群聊数组
	Groups []string `form:"groups" json:"groups"`
	// 正文
	Content string `form:"content" json:"content"`
	// 文件名称
	Filename string `form:"filename" json:"filename"`
}

// SendVideoBatchHandle 群发视频消息
func SendVideoBatchHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendTextBatchRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	self, _ := bot.GetCurrentUser()
	filename := path.Base(res.Content)
	destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
	file, err := utils.FetchFile(res.Content, destPath, filename)
	if err != nil {
		core.FailWithMessage("拉取视频失败"+err.Error(), ctx)
		return
	}
	defer os.Remove(destPath + filename)
	//好友
	for _, item := range res.Friends {
		friend, _ := FindFriend(bot, item, ctx)
		if friend != nil {
			friend.SendVideo(file)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}
	file.Seek(0, 0)
	//群聊
	for _, item := range res.Groups {
		group, _ := FindGroup(bot, item, ctx)
		if group != nil {
			group.SendVideo(file)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}

	core.Ok(ctx)
}

// SendFileBatchHandle 群发文件消息
func SendFileBatchHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendFileBatchRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	self, _ := bot.GetCurrentUser()
	var filename string
	if len(res.Filename) > 0 {
		filename = res.Filename
	} else {
		filename = path.Base(res.Content)
	}
	destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
	file, err := utils.FetchFile(res.Content, destPath, filename)
	if err != nil {
		core.FailWithMessage("拉取文件失败"+err.Error(), ctx)
		return
	}
	defer os.Remove(destPath + filename)
	//好友
	if len(res.Friends) > 0 {
		var friends = make(openwechat.Friends, 0)
		var delays = []time.Duration{}
		for _, item := range res.Friends {
			friend, _ := FindFriend(bot, item, ctx)
			var delay time.Duration
			if friend != nil {
				friends = append(friends, friend)
				rand.Seed(time.Now().UnixNano())
				delay = time.Duration((rand.Intn(3)+1)*1000) * time.Millisecond
				delays = append(delays, delay)
			}
			if friends.Count() > 0 {
				if err := friends.SendFile(file, delays...); err != nil {
					core.FailWithMessage("群发好友出错："+err.Error(), ctx)
					return
				}
			}
		}
	}
	//群聊
	if len(res.Groups) > 0 {
		var groups = make(openwechat.Groups, 0)
		var delays = []time.Duration{}
		for _, item := range res.Groups {
			group, _ := FindGroup(bot, item, ctx)
			var delay time.Duration
			if group != nil {
				groups = append(groups, group)
				rand.Seed(time.Now().UnixNano())
				delay = time.Duration((rand.Intn(3)+1)*1000) * time.Millisecond
				delays = append(delays, delay)
			}
			if groups.Count() > 0 {
				/* file.Seek(0, 0)
				if err := groups.SendFile(file, delays...); err != nil {
					core.FailWithMessage("群发群聊出错："+err.Error(), ctx)
					return
				} */ //源码漏了
			}
		}
	}
	core.Ok(ctx)
}

// SendImgBatchHandle 群发图片消息
func SendImgBatchHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendTextBatchRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	self, _ := bot.GetCurrentUser()
	filename := path.Base(res.Content)
	destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
	file, err := utils.FetchFile(res.Content, destPath, filename)
	if err != nil {
		core.FailWithMessage("拉取图片失败"+err.Error(), ctx)
		return
	}
	defer os.Remove(destPath + filename)
	//好友
	if len(res.Friends) > 0 {
		var friends = make(openwechat.Friends, 0)
		var delays = []time.Duration{}
		for _, item := range res.Friends {
			friend, _ := FindFriend(bot, item, ctx)
			var delay time.Duration
			if friend != nil {
				friends = append(friends, friend)
				rand.Seed(time.Now().UnixNano())
				delay = time.Duration((rand.Intn(3)+1)*1000) * time.Millisecond
				delays = append(delays, delay)
			}
			if friends.Count() > 0 {
				if err := friends.SendImage(file, delays...); err != nil {
					core.FailWithMessage("群发好友出错："+err.Error(), ctx)
					return
				}
			}
		}
	}
	//群聊
	if len(res.Groups) > 0 {
		var groups = make(openwechat.Groups, 0)
		var delays = []time.Duration{}
		for _, item := range res.Groups {
			group, _ := FindGroup(bot, item, ctx)
			var delay time.Duration
			if group != nil {
				groups = append(groups, group)
				rand.Seed(time.Now().UnixNano())
				delay = time.Duration((rand.Intn(3)+1)*1000) * time.Millisecond
				delays = append(delays, delay)
			}
			if groups.Count() > 0 {
				file.Seek(0, 0)
				if err := groups.SendImage(file, delays...); err != nil {
					core.FailWithMessage("群发群聊出错："+err.Error(), ctx)
					return
				}
			}
		}
	}
	core.Ok(ctx)
}

// SendTextBatchHandle 群发文本消息
func SendTextBatchHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendTextBatchRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	//好友
	if len(res.Friends) > 0 {
		var friends = make(openwechat.Friends, 0)
		var delays = []time.Duration{}
		for _, item := range res.Friends {
			friend, _ := FindFriend(bot, item, ctx)
			var delay time.Duration
			if friend != nil {
				friends = append(friends, friend)
				rand.Seed(time.Now().UnixNano())
				delay = time.Duration((rand.Intn(3)+1)*1000) * time.Millisecond
				delays = append(delays, delay)
			}
			if friends.Count() > 0 {
				if err := friends.SendText(res.Content, delays...); err != nil {
					core.FailWithMessage("群发好友出错："+err.Error(), ctx)
					return
				}
			}
		}
	}
	//群聊
	if len(res.Groups) > 0 {
		var groups = make(openwechat.Groups, 0)
		var delays = []time.Duration{}
		for _, item := range res.Groups {
			group, _ := FindGroup(bot, item, ctx)
			var delay time.Duration
			if group != nil {
				groups = append(groups, group)
				rand.Seed(time.Now().UnixNano())
				delay = time.Duration((rand.Intn(3)+1)*1000) * time.Millisecond
				delays = append(delays, delay)
			}
			if groups.Count() > 0 {
				if err := groups.SendText(res.Content, delays...); err != nil {
					core.FailWithMessage("群发群聊出错："+err.Error(), ctx)
					return
				}
			}
		}
	}
	core.Ok(ctx)
}

// SendVideoToGroup 向指定群聊发视频
func SendVideoToGroupHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendFileRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	group, self := FindGroup(bot, res.To, ctx)
	if group == nil {
		return
	}

	var filename string
	if len(res.Filename) > 0 {
		filename = res.Filename
	} else {
		filename = path.Base(res.Content)
	}

	destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
	file, err := utils.FetchFile(res.Content, destPath, filename)
	if err != nil {
		core.FailWithMessage("拉取视频失败"+err.Error(), ctx)
		return
	}
	defer os.Remove(destPath + filename)

	// 发送消息
	if _, err := group.SendVideo(file); err != nil {
		fmt.Println(self.Uin)
		core.FailWithMessage("发送视频失败"+err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

// SendVideoToFriendHandle 向指定用户发video
func SendVideoToFriendHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendFileRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	friend, self := FindFriend(bot, res.To, ctx)
	if friend == nil {
		return
	}

	var filename string
	if len(res.Filename) > 0 {
		filename = res.Filename
	} else {
		filename = path.Base(res.Content)
	}
	destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
	file, err := utils.FetchFile(res.Content, destPath, filename)
	if err != nil {
		core.FailWithMessage("拉取文件失败"+err.Error(), ctx)
		return
	}
	defer os.Remove(destPath + filename)

	// 发送消息
	if _, err := friend.SendVideo(file); err != nil {
		fmt.Println(self.Uin)
		core.FailWithMessage("发送视频失败"+err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

// SendFileToGroup 向指定群聊发文件
func SendFileToGroupHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendFileRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	group, self := FindGroup(bot, res.To, ctx)
	if group == nil {
		return
	}

	var filename string
	if len(res.Filename) > 0 {
		filename = res.Filename
	} else {
		filename = path.Base(res.Content)
	}

	destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
	file, err := utils.FetchFile(res.Content, destPath, filename)
	if err != nil {
		core.FailWithMessage("拉取图片失败"+err.Error(), ctx)
		return
	}
	defer os.Remove(destPath + filename)

	// 发送消息
	if _, err := group.SendFile(file); err != nil {
		fmt.Println(self.Uin)
		core.FailWithMessage("发送文件失败"+err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

// SendFileToFriendHandle 向指定用户发文件
func SendFileToFriendHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendFileRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	friend, self := FindFriend(bot, res.To, ctx)
	if friend == nil {
		return
	}

	var filename string
	if len(res.Filename) > 0 {
		filename = res.Filename
	} else {
		filename = path.Base(res.Content)
	}
	destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
	file, err := utils.FetchFile(res.Content, destPath, filename)
	if err != nil {
		core.FailWithMessage("拉取图片失败"+err.Error(), ctx)
		return
	}
	defer os.Remove(destPath + filename)

	// 发送消息
	if _, err := friend.SendFile(file); err != nil {
		fmt.Println(self.Uin)
		core.FailWithMessage("发送文件失败"+err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

// SendImgToGroupHandle 向指定群聊发图片
func SendImgToGroupHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendMsgRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	group, self := FindGroup(bot, res.To, ctx)
	if group == nil {
		return
	}

	filename := path.Base(res.Content)
	destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
	file, err := utils.FetchFile(res.Content, destPath, filename)
	if err != nil {
		core.FailWithMessage("拉取图片失败"+err.Error(), ctx)
		return
	}
	defer os.Remove(destPath + filename)

	// 发送消息
	if _, err := group.SendImage(file); err != nil {
		fmt.Println(self.Uin)
		core.FailWithMessage("发送图片失败"+err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

// SendImgToFriend 向指定用户发图片
func SendImgToFriendHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendMsgRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	friend, self := FindFriend(bot, res.To, ctx)
	if friend == nil {
		return
	}

	filename := path.Base(res.Content)
	destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
	file, err := utils.FetchFile(res.Content, destPath, filename)
	if err != nil {
		core.FailWithMessage("拉取图片失败"+err.Error(), ctx)
		return
	}
	defer os.Remove(destPath + filename)

	// 发送消息
	if _, err := friend.SendImage(file); err != nil {
		fmt.Println(self.Uin)
		core.FailWithMessage("发送图片失败"+err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

// SendTextToFriend 向指定用户发消息
func SendTextToFriendHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendMsgRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")

	bot := global.GetBot(appKey)
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
	// 发送消息
	if _, err := friend.SendText(res.Content); err != nil {
		core.FailWithMessage(err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

// SendTextToGroup 向指定群组发送消息
func SendTextToGroupHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendMsgRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")

	bot := global.GetBot(appKey)
	// 获取登录用户
	self, _ := bot.GetCurrentUser()
	// 获取所有群组
	groups, err := self.Groups(true)
	if err != nil {
		core.FailWithMessage("群组获取失败", ctx)
		return
	}
	// 判断指定群组是否存在
	search := groups.SearchByUserName(1, res.To)
	if search.Count() < 1 {
		core.FailWithMessage("指定群组不存在", ctx)
		return
	}
	// 取出指定群组
	group := search.First()
	// 发送消息
	if _, err := group.SendText(res.Content); err != nil {
		core.FailWithMessage(err.Error(), ctx)
		return
	}
	core.Ok(ctx)
}

//FindFriend 根据username获取好友
func FindFriend(bot *protocol.WechatBot, username string, ctx *gin.Context) (*openwechat.Friend, *openwechat.Self) {
	// 获取登录用户
	self, _ := bot.GetCurrentUser()
	// 查找指定的好友
	friends, _ := self.Friends(true)
	// 查询指定好友
	friendSearchResult := friends.SearchByUserName(1, username)
	//friendSearchResult := friends.SearchByNickName(1, username)
	if friendSearchResult.Count() < 1 {
		core.FailWithMessage("指定好友不存在", ctx)
		return nil, self
	}
	// 取出好友
	return friendSearchResult.First(), self
}

//FindGroup 根据username获取群组
func FindGroup(bot *protocol.WechatBot, username string, ctx *gin.Context) (*openwechat.Group, *openwechat.Self) {
	// 获取登录用户
	self, _ := bot.GetCurrentUser()
	// 查找指定的好友
	groups, _ := self.Groups(true)
	// 查询指定好友
	searchResult := groups.SearchByUserName(1, username)
	if searchResult.Count() < 1 {
		core.FailWithMessage("指定群组不存在", ctx)
		return nil, self
	}
	// 取出数据
	return searchResult.First(), self
}
