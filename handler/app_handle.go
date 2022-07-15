package handler

import (
	"encoding/xml"
	"go-wxbot/global"
	"go-wxbot/logger"

	"github.com/eatmoreapple/openwechat"
)

type AppMessageData struct {
	XMLName xml.Name `xml:"msg"`
	Appmsg  struct {
		Appid             string `xml:"appid,attr"`
		Sdkver            string `xml:"sdkver,attr"`
		Title             string `xml:"title"`
		Des               string `xml:"des"`
		Action            string `xml:"action"`
		Type              string `xml:"type"`
		Showtype          string `xml:"showtype"`
		Content           string `xml:"content"`
		URL               string `xml:"url"`
		Dataurl           string `xml:"dataurl"`
		Lowurl            string `xml:"lowurl"`
		Lowdataurl        string `xml:"lowdataurl"`
		Recorditem        string `xml:"recorditem"`
		Thumburl          string `xml:"thumburl"`
		Messageaction     string `xml:"messageaction"`
		Md5               string `xml:"md5"`
		Extinfo           string `xml:"extinfo"`
		Sourceusername    string `xml:"sourceusername"`
		Sourcedisplayname string `xml:"sourcedisplayname"`
		Commenturl        string `xml:"commenturl"`
		Appattach         struct {
			Totallen          string `xml:"totallen"`
			Attachid          string `xml:"attachid"`
			Emoticonmd5       string `xml:"emoticonmd5"`
			Fileext           string `xml:"fileext"`
			Fileuploadtoken   string `xml:"fileuploadtoken"`
			OverwriteNewmsgid string `xml:"overwrite_newmsgid"`
			Filekey           string `xml:"filekey"`
			Cdnattachurl      string `xml:"cdnattachurl"`
			Aeskey            string `xml:"aeskey"`
			Encryver          string `xml:"encryver"`
		} `xml:"appattach"`
		Weappinfo struct {
			Pagepath       string `xml:"pagepath"`
			Username       string `xml:"username"`
			Appid          string `xml:"appid"`
			Appservicetype string `xml:"appservicetype"`
		} `xml:"weappinfo"`
		Websearch string `xml:"websearch"`
	} `xml:"appmsg"`
	Fromusername string `xml:"fromusername"`
	Scene        string `xml:"scene"`
	Appinfo      struct {
		Version string `xml:"version"`
		Appname string `xml:"appname"`
	} `xml:"appinfo"`
	Commenturl string `xml:"commenturl"`
}

// APP消息处理
func appMessageHandle(ctx *openwechat.MessageContext) {
	// 取出发送者
	sender, _ := ctx.Sender()
	senderUser := sender.NickName

	logger.Log.Infof("[收到小程序消息]\n消息类型: %v\n 发信人：%v\n 内容：%v", ctx.MsgType, senderUser, ctx.Content)
	msg := ctx.Message
	bot := ctx.Bot
	var resp = CallbackRes{Type: global.MSG_APP, MsgId: msg.MsgId, From: sender.UserName, NickName: sender.NickName, Content: msg.Content}

	if !ctx.IsSendBySelf() {
		if ctx.IsSendByGroup() {
			// 取出消息在群里面的发送者
			senderInGroup, _ := ctx.SenderInGroup()
			resp.Useringroup = senderInGroup.NickName + senderInGroup.UserName
		}
	}

	NotifyWebhook(bot, &resp)
	ctx.Next()
}
