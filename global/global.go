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
	MSG_APP         int32 = 2002  // 小程序消息
	MSG_GROUPINVITE int32 = 2003  //群邀请
	MSG_RECEIVEFILE int32 = 2004  //接收文件
	MSG_VERIFY      int32 = 37    // 好友验证
	MSG_SHARECARD   int32 = 42    // 名片消息
	MSG_SYS         int32 = 10000 // 系统消息
	MSG_RECALLED    int32 = 2005  // 消息撤回

	EVENT_PRIVATE_CHAT      string = "EventPrivateChat"
	EVENT_GROUP_CHAT        string = "EventGroupChat"
	EVENT_FRIEND_VERIFY     string = "EventFriendVerify"
	EVENT_GROUP_MEMBER_ADD  string = "EventGroupMemberAdd"
	EVENT_GROUP_MEMBER_DECR string = "EventGroupMemberDecrease"
	EVENT_GROUP_NAME_CHANGE string = "EventGroupNameChange"  //群名修改
	EVENT_INVITEED_IN_GROUP string = "EventInvitedInGroup"   //被邀请入群事
	EVENT_RECEIVED_TRANSFER string = "EventReceivedTransfer" //收到转账事件（收到好友转账时，运行这里）
	EVENT_SCAN_CACH_MONEY   string = "EventScanCashMoney"    //面对面收款（二维码收款时，运行这里）
	EVENT_SYS_MSG           string = "EventSysMsg"           //系统消息事件
)
