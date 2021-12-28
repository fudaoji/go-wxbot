package handler

import (
	"go-wxbot/logger"

	"github.com/eatmoreapple/openwechat"
)

const (
	MSGTYPE_TEXT string = "text"
)

// 处理文本消息
func textMessageHandle(ctx *openwechat.MessageContext) {
	sender, _ := ctx.Sender()
	senderUser := sender.NickName

	logger.Log.Infof("[收到新文字消息] == 发信人：%v ==> 内容：%v", senderUser, ctx.Content)
	msg := ctx.Message
	bot := ctx.Bot
	var resp = CallbackRes{From: sender.UserName, Type: MSGTYPE_TEXT, Content: msg.Content}

	if !ctx.IsSendBySelf() {
		if ctx.IsSendByGroup() {
			// 取出消息在群里面的发送者
			senderInGroup, _ := ctx.SenderInGroup()
			resp.Useringroup = senderInGroup.NickName + senderInGroup.UserName
		}
	}

	NotifyWebhook(bot, &resp)

	//if !ctx.IsSendBySelf() {
	//	sender, _ := ctx.Sender()
	//	if ctx.IsSendByGroup() {
	//		// 取出消息在群里面的发送者
	//		senderInGroup, _ := ctx.SenderInGroup()
	//		logger.Log.Infof("[群聊][收到新文字消息] == 发信人：%v[%v] ==> 内容：%v", sender.NickName,
	//			senderInGroup.NickName, ctx.Content)
	//	} else {
	//		logger.Log.Infof("[好友][收到新文字消息] == 发信人：%v ==> 内容：%v", sender.NickName, ctx.Content)
	//	}
	//}
	ctx.Next()
}
