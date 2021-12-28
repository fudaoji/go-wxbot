package handler

import (
	"go-wxbot/logger"

	"github.com/eatmoreapple/openwechat"
)

func checkIsCanRead(message *openwechat.Message) bool {
	// 通知消息和自己发的不处理
	return !message.IsNotify() && !message.IsSendBySelf()
}

// 设置消息为已读
func setTheMessageAsRead(ctx *openwechat.MessageContext) {
	err := ctx.AsRead()
	if err != nil {
		logger.Log.Errorf("设置消息为已读出错: %v", err)
	}
	ctx.Next()
}
