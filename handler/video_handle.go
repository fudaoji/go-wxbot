package handler

import (
	"go-wxbot/global"
	"go-wxbot/logger"

	"github.com/eatmoreapple/openwechat"
)

// 处理语音消息
func videoMessageHandle(ctx *openwechat.MessageContext) {
	sender, _ := ctx.Sender()
	senderUser := sender.NickName

	logger.Log.Infof("[收到新视频消息] == 发信人：%v ==> 内容：%v", senderUser, ctx.Content)
	msg := ctx.Message
	bot := ctx.Bot
	var resp = CallbackRes{Type: global.MSG_VIDEO, MsgId: msg.MsgId, From: sender.UserName, NickName: sender.NickName, Content: msg.Content}

	if !ctx.IsSendBySelf() {
		if ctx.IsSendByGroup() {
			resp.Event = global.EVENT_GROUP_CHAT
			// 取出消息在群里面的发送者
			senderInGroup, _ := ctx.SenderInGroup()
			resp.Useringroup = senderInGroup.NickName + senderInGroup.UserName
			resp.Group = sender.UserName
			resp.GroupName = sender.NickName
			resp.From = senderInGroup.UserName
			resp.NickName = senderInGroup.NickName
		} else {
			resp.Event = global.EVENT_PRIVATE_CHAT
		}
	}

	NotifyWebhook(bot, &resp)
	ctx.Next()
}
