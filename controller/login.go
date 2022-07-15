package controller

import (
	"go-wxbot/core"
	"go-wxbot/global"
	"go-wxbot/handler"
	"go-wxbot/logger"
	"go-wxbot/model"
	"go-wxbot/protocol"

	"github.com/eatmoreapple/openwechat"
	"github.com/gin-gonic/gin"
)

// 获取登录URL返回结构体
type loginUrlResponse struct {
	Uuid string `json:"uuid"`
	Url  string `json:"url"`
}

// 登录请求体
type loginRes struct {
	// 回调地址
	Webhook string `form:"webhook" json:"webhook"`
}

// GetLoginUrlHandle 获取登录扫码连接
func GetLoginUrlHandle(ctx *gin.Context) {
	appKey := ctx.Request.Header.Get("AppKey")

	// 获取一个微信机器人对象
	bot := handler.InitWechatBotHandle()

	// 获取登录二维码链接
	url := "https://login.weixin.qq.com/qrcode/"
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl
	uuid, err := bot.GetUUID()
	if err != nil {
		core.FailWithMessage("获取UUID失败", ctx)
		return
	}
	// 拼接URL
	url = url + *uuid

	// 保存Bot到实例
	global.SetBot(appKey, bot)

	// 返回数据
	core.OkWithData(loginUrlResponse{Uuid: *uuid, Url: url}, ctx)
}

// LoginHandle 登录
func LoginHandle(ctx *gin.Context) {
	appKey := ctx.Request.Header.Get("AppKey")
	uuid := ctx.Query("uuid")
	if len(uuid) < 1 {
		core.FailWithMessage("uuid为必传参数", ctx)
		return
	}
	// 获取Bot对象
	bot := global.GetBot(appKey)
	if bot == nil {
		bot = handler.InitWechatBotHandle()
		global.SetBot(appKey, bot)
	}

	// 已扫码回调
	bot.ScanCallBack = func(body []byte) {
		logger.Log.Infof("[%v]已扫码", appKey)
	}

	// 设置登录成功回调
	bot.LoginCallBack = func(body []byte) {
		logger.Log.Infof("[%v]登录成功", appKey)
	}

	// 热登录
	storage := protocol.NewRedisHotReloadStorage("wechat:login:" + appKey)
	if err := bot.HotLoginWithUUID(uuid, storage, true); err != nil {
		logger.Log.Errorf("热登录失败: %v", err)
		core.FailWithMessage("登录失败："+err.Error(), ctx)
		return
	}

	user, err := bot.GetCurrentUser()
	if err != nil {
		logger.Log.Errorf("获取登录用户信息失败: %v", err.Error())
		core.FailWithMessage("获取登录用户信息失败："+err.Error(), ctx)
		return
	} else {
		//webhook和uin入库
		var resData loginRes
		err := ctx.ShouldBindJSON(&resData)
		if err == nil && len(resData.Webhook) > 0 {
			logger.Log.Infof("保存webhook[%v]", resData.Webhook)
			model.UpdateWebhookByAppkey(appKey, resData.Webhook, user.Uin)
		}
	}
	logger.Log.Infof("当前登录用户：%v", user.NickName)
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
	//core.OkWithMessage("登录成功", ctx)
}
