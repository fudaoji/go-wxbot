package global

import (
	"go-wxbot/protocol"
)

var (
	// 登录用户的Bot对象
	wechatBots map[string]*protocol.WechatBot
)

const (
	MSGTYPE_TEXT           string = "text"           // 文本消息
	MSGTYPE_IMAGE          string = "image"          // 图片消息
	MSGTYPE_VOICE          string = "voice"          // 语音消息
	MSGTYPE_VIDEO          string = "video"          // 视频消息
	MSGTYPE_EMOTICON       string = "emotion"        // 表情消息
	MSGTYPE_VERIFY         string = "verify"         // 认证消息
	MSGTYPE_POSSIBLEFRIEND string = "possiblefriend" // 好友推荐
	MSGTYPE_SHARECARD      string = "sharecard"      // 名片消息

	MSGTYPE_LOCATION   string = "location"   // 地理位置消息
	MSGTYPE_APP        string = "app"        // APP消息
	MSGTYPE_VOIP       string = "voip"       // VOIP消息
	MSGTYPE_VOIPNOTIFY string = "voipnotify" // voip结束消息
	MSGTYPE_VOIPINVITE string = "voipinvite" // VOIP邀请
	MSGTYPE_MICROVIDEO string = "microvideo" // 小视频消息
	MSGTYPE_SYS        string = "sys"        // 系统消息
	MSGTYPE_RECALLED   string = "recalled"   // 消息撤回
)
