package controller

import (
	"fmt"
	"go-wxbot/core"
	"go-wxbot/global"
	"go-wxbot/protocol"
	"os"
	"path"
	"strings"

	"github.com/fudaoji/go-utils"

	"github.com/eatmoreapple/openwechat"
	"github.com/gin-gonic/gin"
)

// 发送消息请求体
type sendMsgRes struct {
	// username
	To string `form:"to" json:"to"`
	// 正文
	Content string `form:"msg" json:"msg"`
}

// 发送文件请求体
type sendFileRes struct {
	// username
	To string `form:"to" json:"to"`
	// 文件路径
	Content string `form:"path" json:"path"`
	// 文件名称
	Filename string `form:"filename" json:"filename"`
}

// SendFileHandle 发文件
func SendFileHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendFileRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	member, self := FindMember(bot, res.To, ctx)
	if member == nil {
		return
	}

	var file *os.File
	var err error
	if strings.Contains(res.Content, "http") {
		var filename string
		if len(res.Filename) > 0 {
			filename = res.Filename
		} else {
			filename = path.Base(res.Content)
		}

		destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
		file, err = utils.FetchFile(res.Content, destPath, filename)
		if err != nil {
			core.FailWithMessage("拉取文件失败"+err.Error(), ctx)
			return
		}
		defer os.Remove(destPath + filename)
	} else {
		if file, err = os.Open(res.Content); err != nil {
			core.FailWithMessage("打开文件失败:"+err.Error(), ctx)
			return
		}
		defer file.Close()
	}

	// 发送消息
	if member.IsFriend() {
		if _, err := self.SendFileToFriend(&openwechat.Friend{member}, file); err != nil {
			core.FailWithMessage("发送文件失败"+err.Error(), ctx)
			return
		}
	}
	if member.IsGroup() {
		if _, err := self.SendFileToGroup(&openwechat.Group{member}, file); err != nil {
			core.FailWithMessage("发送文件失败"+err.Error(), ctx)
			return
		}
	}
	if member.IsMP() {
		if _, err := self.SendFileToMp(&openwechat.Mp{member}, file); err != nil {
			core.FailWithMessage("发送文件失败"+err.Error(), ctx)
			return
		}
	}
	core.Ok(ctx)
}

// SendVideoHandle 向指定用户发video
func SendVideoHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendFileRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	member, self := FindMember(bot, res.To, ctx)
	if member == nil {
		return
	}

	var file *os.File
	var err error
	if strings.Contains(res.Content, "http") {
		var filename string
		if len(res.Filename) > 0 {
			filename = res.Filename
		} else {
			filename = path.Base(res.Content)
		}
		destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
		file, err = utils.FetchFile(res.Content, destPath, filename)
		if err != nil {
			core.FailWithMessage("拉取文件失败"+err.Error(), ctx)
			return
		}
		defer os.Remove(destPath + filename)
	} else {
		if file, err = os.Open(res.Content); err != nil {
			core.FailWithMessage("打开文件失败:"+err.Error(), ctx)
			return
		}
		defer file.Close()
	}

	// 发送消息
	if member.IsFriend() {
		if _, err := self.SendVideoToFriend(&openwechat.Friend{member}, file); err != nil {
			core.FailWithMessage("发送视频失败"+err.Error(), ctx)
			return
		}
	}
	if member.IsGroup() {
		if _, err := self.SendVideoToGroup(&openwechat.Group{member}, file); err != nil {
			core.FailWithMessage("发送视频失败"+err.Error(), ctx)
			return
		}
	}
	core.Ok(ctx)
}

// SendImg 向指定对象发图片
func SendImgHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendMsgRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}

	bot := GetCurBot(ctx)
	member, self := FindMember(bot, res.To, ctx)
	if member == nil {
		return
	}

	var file *os.File
	var err error
	if strings.Contains(res.Content, "http") {
		filename := path.Base(res.Content)
		destPath := fmt.Sprintf("%s%d/", core.GetVal("uploadpath", "./uploads/"), self.Uin)
		file, err = utils.FetchFile(res.Content, destPath, filename)
		if err != nil {
			core.FailWithMessage("拉取图片失败:"+err.Error(), ctx)
			return
		}
		defer os.Remove(destPath + filename)
	} else {
		if file, err = os.Open(res.Content); err != nil {
			core.FailWithMessage("打开图片失败:"+err.Error(), ctx)
			return
		}
		defer file.Close()
	}

	if member.IsFriend() {
		if _, err := self.SendImageToFriend(&openwechat.Friend{member}, file); err != nil {
			core.FailWithMessage("发送图片失败:"+err.Error(), ctx)
			return
		}
	}
	if member.IsGroup() {
		if _, err := self.SendImageToGroup(&openwechat.Group{member}, file); err != nil {
			core.FailWithMessage("发送图片失败:"+err.Error(), ctx)
			return
		}
	}
	if member.IsMP() {
		if _, err := self.SendImageToMp(&openwechat.Mp{member}, file); err != nil {
			core.FailWithMessage("发送图片失败"+err.Error(), ctx)
			return
		}
	}
	core.Ok(ctx)
}

// SendText 发文本消息
func SendTextHandle(ctx *gin.Context) {
	// 取出请求参数
	var res sendMsgRes
	if err := ctx.ShouldBindJSON(&res); err != nil {
		core.FailWithMessage("参数获取失败", ctx)
		return
	}
	// 获取AppKey
	appKey := ctx.Request.Header.Get("AppKey")

	bot := global.GetBot(appKey)
	member, self := FindMember(bot, res.To, ctx)
	if member == nil {
		return
	}
	if member.IsFriend() {
		if _, err := self.SendTextToFriend(&openwechat.Friend{member}, res.Content); err != nil {
			core.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	if member.IsGroup() {
		if _, err := self.SendTextToGroup(&openwechat.Group{member}, res.Content); err != nil {
			core.FailWithMessage(err.Error(), ctx)
			return
		}
	}
	if member.IsMP() {
		if _, err := self.SendTextToMp(&openwechat.Mp{member}, res.Content); err != nil {
			core.FailWithMessage(err.Error(), ctx)
			return
		}
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

//FindMember 根据username获取通讯录
func FindMember(bot *protocol.WechatBot, username string, ctx *gin.Context) (*openwechat.User, *openwechat.Self) {
	// 获取登录用户
	self, _ := bot.GetCurrentUser()
	// 查找指定的好友
	members, _ := self.Members(true)
	// 查询指定好友
	friendSearchResult := members.SearchByUserName(1, username)
	if friendSearchResult.Count() < 1 {
		core.FailWithMessage("指定好友不存在", ctx)
		return nil, self
	}
	// 取出好友
	return friendSearchResult.First(), self
}

// SendFileBatchHandle 群发文件消息
/*func SendFileBatchHandle(ctx *gin.Context) {
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
		for i, item := range res.Groups {
			group, _ := FindGroup(bot, item, ctx)
			if group != nil {
				group.SendFile(file)
				file.Seek(0, 0)
				if i < len(res.Groups)-1 {
					time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
				}
			}
		}
	}
	core.Ok(ctx)
}*/
