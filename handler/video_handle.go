package handler

import (
	"bytes"
	"fmt"
	"go-wxbot/core"
	"go-wxbot/global"
	"go-wxbot/logger"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/eatmoreapple/openwechat"
	"github.com/fudaoji/go-utils"
)

//checkIsFriendAdd 判断是否系统消息
func checkIsVideo(message *openwechat.Message) bool {
	return message.IsVideo()
}

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

	// 下载视频资源
	fileResp, err := ctx.GetFile()
	if err != nil {
		logger.Log.Errorf("文件下载失败: %v", err.Error())
		return
	}
	defer fileResp.Body.Close()
	imgFileByte, err := ioutil.ReadAll(fileResp.Body)
	if err != nil {
		logger.Log.Errorf("文件读取错误: %v", err.Error())
		return
	} else {
		// 读取文件相关信息
		contentType := http.DetectContentType(imgFileByte)
		fileType := strings.Split(contentType, "/")[1]
		fileName := fmt.Sprintf("%v.%v", ctx.MsgId, fileType)
		path := core.GetVal("uploadpath", "./uploads/")
		if user, err := ctx.Bot.GetCurrentUser(); err == nil {
			path = fmt.Sprintf("%v/%v/", path, user.Uin)
		}

		// 保存文件
		reader2 := ioutil.NopCloser(bytes.NewReader(imgFileByte))
		_, err := utils.SaveFile(reader2, path, fileName)
		if err != nil {
			logger.Log.Errorf("保存文件错误: %v", err.Error())
			return
		}

		resp.Content = path + fileName
	}

	NotifyWebhook(bot, &resp)
	ctx.Next()
}
