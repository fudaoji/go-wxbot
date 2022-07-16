package handler

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go-wxbot/core"
	"go-wxbot/global"
	"go-wxbot/logger"
	"io/ioutil"

	"github.com/eatmoreapple/openwechat"
	"github.com/fudaoji/go-utils"
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
	if ctx.IsComeFromGroup() {
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

	// 解析文件
	var data AppMessageData
	if err := xml.Unmarshal([]byte(ctx.Content), &data); err != nil {
		logger.Log.Errorf("消息解析失败: %v", err.Error())
		logger.Log.Debugf("原始内容: %v", ctx.Content)
		return
	} else {
		tt := data.Appmsg.Type
		dealType := []string{"2", "3", "4", "6", "15"}
		if !checkIsExist(dealType, tt) {
			logger.Log.Infof("奇奇怪怪的未定义处理类型，跳过处理。类型: %v", tt)
			return
		}
		// 下载文件资源
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
			//contentType := http.DetectContentType(imgFileByte)
			fileName := fmt.Sprintf("%v_%v", ctx.MsgId, data.Appmsg.Title)
			path := core.GetVal("uploadpath", "./uploads/")
			if user, err := ctx.Bot.GetCurrentUser(); err == nil {
				path = fmt.Sprintf("%v/%v/", path, user.Uin)
			}
			// 上传文件(reader2解决上传空文件的BUG,因为http.Response.Body只允许读一次)
			reader2 := ioutil.NopCloser(bytes.NewReader(imgFileByte))

			_, err := utils.SaveFile(reader2, path, fileName)
			if err != nil {
				logger.Log.Errorf("保存文件错误: %v", err.Error())
				return
			}

			resp.Content = path + fileName
		}
	}

	NotifyWebhook(bot, &resp)
	ctx.Next()
}

// 判断数组是否包含某个元素
func checkIsExist(a []string, k string) bool {
	for _, item := range a {
		if item == k {
			return true
		}
	}
	return false
}
