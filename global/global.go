package global

import (
	"go-wxbot/protocol"
)

var (
	// 登录用户的Bot对象
	wechatBots map[string]*protocol.WechatBot
)

/**
*消息类型
● 文本    1
● 图片消息    3
● 语音消息    34
● 好友验证    37
● 名片推荐    42
● 视频消息    43
● 动态表情    47
● 位置消息    48
● 分享链接    49
● 转账消息    2000
● 红包消息    2001
● 小程序    2002
● 群邀请    2003
● 接收文件    2004
● 撤回消息    2005
● 系统消息    10000
● 服务通知    10002
*/
const (
	MSG_TEXT        int32 = 1     // 文本消息
	MSG_IMAGE       int32 = 3     // 图片消息
	MSG_VOICE       int32 = 34    // 语音消息
	MSG_VIDEO       int32 = 43    // 视频消息
	MSG_EMOTICON    int32 = 47    // 表情消息
	MSG_LOCATION    int32 = 48    // 地理位置消息
	MSG_LINK        int32 = 49    // 分享链接消息
	MSG_TRANSFER    int32 = 2000  //转账消息
	MSG_RED         int32 = 2001  //红包消息
	MSG_MINIAPP     int32 = 2002  // 小程序消息
	MSG_GROUPINVITE int32 = 2003  //群邀请
	MSG_RECEIVEFILE int32 = 2004  //接收文件
	MSG_VERIFY      int32 = 37    // 好友验证
	MSG_SHARECARD   int32 = 42    // 名片消息
	MSG_SYS         int32 = 10000 // 系统消息
	MSG_RECALLED    int32 = 2005  // 消息撤回

	MSGTYPE_TEXT           string = "text"           // 文本消息
	MSGTYPE_IMAGE          string = "image"          // 图片消息
	MSGTYPE_VOICE          string = "voice"          // 语音消息
	MSGTYPE_VIDEO          string = "video"          // 视频消息
	MSGTYPE_EMOTICON       string = "emotion"        // 表情消息
	MSGTYPE_VERIFY         string = "verify"         // 认证消息
	MSGTYPE_POSSIBLEFRIEND string = "possiblefriend" // 好友推荐
	MSGTYPE_SHARECARD      string = "sharecard"      // 名片消息
	MSGTYPE_LOCATION       string = "location"       // 地理位置消息
	MSGTYPE_APP            string = "app"            // APP消息
	MSGTYPE_VOIP           string = "voip"           // VOIP消息
	MSGTYPE_VOIPNOTIFY     string = "voipnotify"     // voip结束消息
	MSGTYPE_VOIPINVITE     string = "voipinvite"     // VOIP邀请
	MSGTYPE_MICROVIDEO     string = "microvideo"     // 小视频消息
	MSGTYPE_SYS            string = "sys"            // 系统消息
	MSGTYPE_RECALLED       string = "recalled"       // 消息撤回
)
