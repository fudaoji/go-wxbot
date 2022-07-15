package handler

import (
	"fmt"
	. "go-wxbot/core"
	. "go-wxbot/db"
	"go-wxbot/global"
	"go-wxbot/logger"
	"go-wxbot/model"
	"go-wxbot/protocol"

	"github.com/eatmoreapple/openwechat"
)

// InitBotWithStart 系统启动的时候从Redis加载登录信息自动登录
func InitBotWithStart() {
	keys, err := RedisClient.GetKeys("wechat:login:*")
	if err != nil {
		logger.Log.Error("获取Key失败")
		return
	}
	logger.Log.Infof("获取到登录用户信息数量：%v", len(keys))
	for _, key := range keys {
		// 提取出AppKey
		appKey := key[13:]
		// 调用热登录
		logger.Log.Debugf("当前热登录AppKey: %v", appKey)
		bot := InitWechatBotHandle()
		storage := protocol.NewRedisHotReloadStorage(key)
		if err := bot.HotLogin(storage, false); err != nil {
			logger.Log.Infof("[%v] 热登录失败，错误信息：%v", appKey, err.Error())
			// 登录失败，删除热登录数据
			if err := RedisClient.Del(key); err != nil {
				logger.Log.Errorf("[%v] Redis缓存删除失败，错误信息：%v", key, err.Error())
			}
			continue
		}
		loginUser, _ := bot.GetCurrentUser()
		logger.Log.Infof("[%v]初始化自动登录成功，用户名：%v", appKey, loginUser.NickName)
		// 登录成功，写入到WechatBots
		global.SetBot(appKey, bot)
	}
}

// InitWechatBotHandle 初始化微信机器人
func InitWechatBotHandle() *protocol.WechatBot {
	bot := openwechat.DefaultBot(openwechat.Desktop)

	if GetIntVal("app_debug", 1) == 0 {
		bot.SyncCheckCallback = nil
	}
	// 定义读取消息错误回调函数
	//var getMessageErrorCount int32
	//bot.GetMessageErrorHandler = func(err error) {
	//	atomic.AddInt32(&getMessageErrorCount, 1)
	//	// 如果发生了三次错误,那么直接退出
	//	if getMessageErrorCount == 3 {
	//		logger.Log.Errorf("获取消息发生错误达到三次，直接退出。错误信息：%v", err.Error())
	//		_ = bot.Logout()
	//	}
	//}
	// 注册消息处理函数
	HandleMessage(bot)
	// 获取消息发生错误
	//bot.MessageOnError()
	// 返回机器人对象
	return &protocol.WechatBot{Bot: *bot}
}

//HandleMessage  消息处理器
func HandleMessage(bot *openwechat.Bot) {
	// 定义一个处理器
	dispatcher := openwechat.NewMessageMatchDispatcher()
	// 设置为异步处理
	dispatcher.SetAsync(true)
	// 处理消息为已读
	dispatcher.RegisterHandler(checkIsCanRead, setTheMessageAsRead)
	// 未定义消息处理
	dispatcher.RegisterHandler(checkIsOther, otherMessageHandle)

	// 注册文本消息处理函数
	dispatcher.OnText(textMessageHandle)
	// 注册图片消息处理器
	dispatcher.OnImage(imageMessageHandle)
	// 注册语音消息处理函数
	dispatcher.OnVoice(voiceMessageHandle)
	// 注册表情包消息处理器
	dispatcher.OnEmoticon(emoticonMessageHandle)
	// APP消息处理
	dispatcher.OnMedia(appMessageHandle)
	// 保存消息
	//dispatcher.RegisterHandler(checkNeedSave, saveToDb)

	// 注册消息处理器
	bot.MessageHandler = openwechat.DispatchMessage(dispatcher)
}

//NotifyWebhook  通知客户端平台
func NotifyWebhook(bot *openwechat.Bot, data *CallbackRes) {
	user, _ := bot.GetCurrentUser()
	data.Robot = user.Uin
	appkeyRecord := model.Appkey{Uin: user.Uin}
	appkeyRecord.FindByUin()
	if len(appkeyRecord.Webhook) > 0 {
		url := appkeyRecord.Webhook
		data.Appkey = appkeyRecord.Appkey
		ReqPostJson(url, data, nil)
	} else {
		fmt.Println("未填写webhook")
	}
}
