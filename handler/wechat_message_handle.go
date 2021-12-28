package handler

import (
	"fmt"
	. "go-wxbot/core"
	"go-wxbot/model"

	"github.com/eatmoreapple/openwechat"
)

// 回调请求体
type CallbackRes struct {
	Appkey      string      `json:"appkey"`
	From        string      `json:"from"`
	Type        string      `json:"type"`
	Content     interface{} `json:"content"`
	Useringroup string      `json:"useringroup"`
}

func HandleMessage(bot *openwechat.Bot) {
	// 定义一个处理器
	dispatcher := openwechat.NewMessageMatchDispatcher()
	// 设置为异步处理
	dispatcher.SetAsync(true)
	// 处理消息为已读
	dispatcher.RegisterHandler(checkIsCanRead, setTheMessageAsRead)

	// 注册文本消息处理函数
	dispatcher.OnText(textMessageHandle)
	// 注册图片消息处理器
	dispatcher.OnImage(imageMessageHandle)
	// 注册表情包消息处理器
	dispatcher.OnEmoticon(emoticonMessageHandle)
	// APP消息处理
	dispatcher.OnMedia(appMessageHandle)
	// 保存消息
	//dispatcher.RegisterHandler(checkNeedSave, saveToDb)
	// 未定义消息处理
	dispatcher.RegisterHandler(checkIsOther, otherMessageHandle)

	// 注册消息处理器
	bot.MessageHandler = openwechat.DispatchMessage(dispatcher)
}

//NotifyWebhook  通知客户端平台
func NotifyWebhook(bot *openwechat.Bot, data *CallbackRes) {
	user, _ := bot.GetCurrentUser()
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
