package handler

import (
	"go-wxbot/global"
	"go-wxbot/logger"

	"github.com/eatmoreapple/openwechat"
)

//checkIsFriendAdd 判断是否好友添加请求
func checkIsFriendAdd(message *openwechat.Message) bool {
	return message.IsFriendAdd()
}

// 处理好友添加消息
func friendAddMessageHandle(ctx *openwechat.MessageContext) {
	sender, _ := ctx.Sender()
	msg := ctx.Message
	bot := ctx.Bot

	logger.Log.Infof("[收到添加好友请求]\n发信人：%v\n内容：%v", sender.NickName, ctx.Content)

	var resp = CallbackRes{Type: global.MSG_VERIFY, MsgId: msg.MsgId, From: sender.UserName, NickName: sender.NickName, Content: msg.Content}

	resp.Event = global.EVENT_PRIVATE_CHAT
	msg.Agree()
	NotifyWebhook(bot, &resp)
	ctx.Next()
}
