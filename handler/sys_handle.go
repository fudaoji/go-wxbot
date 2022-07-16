package handler

import (
	"go-wxbot/global"
	"go-wxbot/logger"

	"github.com/eatmoreapple/openwechat"
)

//checkIsFriendAdd 判断是否系统消息
func checkIsSystem(message *openwechat.Message) bool {
	return message.IsSystem()
}

// 处理好友添加消息
func sysMessageHandle(ctx *openwechat.MessageContext) {
	sender, _ := ctx.Sender()
	msg := ctx.Message
	bot := ctx.Bot

	logger.Log.Infof("[收到系统消息]\n发信人：%v\n内容：%v", sender.NickName, ctx.Content)
	logger.Log.Infof("消息体：%v", ctx)

	var resp = CallbackRes{Type: global.MSG_SYS, MsgId: msg.MsgId, From: sender.UserName, NickName: sender.NickName, Content: msg.Content}

	resp.Event = global.EVENT_PRIVATE_CHAT

	NotifyWebhook(bot, &resp)
	ctx.Next()
}
