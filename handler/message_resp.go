package handler

type CallbackRes struct {
	Appkey      string      `json:"appkey"`
	Event       string      `json:"event"`
	Robot       int64       `json:"robot_wxid"`
	Type        int32       `json:"type"`
	From        string      `json:"from_wxid"`
	NickName    string      `json:"from_name"`
	Content     interface{} `json:"msg"`
	Useringroup string      `json:"useringroup"`
	Group       string      `json:"from_group"`
	GroupName   string      `json:"from_group_name"`
	MsgId       string      `json:"msgid"`
}

// 私聊回调请求体
/**
* // HTTP(POST)示例
{
		"appkey": "sdfsffdgss", //appkey
   		"event": "EventPrivateChat", // 事件名
        "robot_wxid": "",  // 机器人账号id
        "type": 1,  // 1/文本消息 3/图片消息 34/语音消息  42/名片消息  43/视频 47/动态表情 48/地理位置  49/分享链接  2001/红包  2002/小程序  2003/群邀请  更多请参考常量表
        "from_wxid": "",  // 来源用户ID
        "from_name": "",  // 来源用户昵称
        "msg": "",  // 消息内容
        "msgid": 0  // 消息ID
}
*/
type PrivateChatResp struct {
	Appkey   string      `json:"appkey"`
	Event    string      `json:"event"`
	Robot    string      `json:"robot_wxid"`
	Type     string      `json:"type"`
	From     string      `json:"from_wxid"`
	NickName string      `json:"from_name"`
	Content  interface{} `json:"msg"`
	MsgId    string      `json:"msgid"`
}

// 群聊回调请求体
/**
* // HTTP(POST)示例
{
		"appkey": "sdfsffdgss", //appkey
   		"event": "EventPrivateChat", // 事件名
        "robot_wxid": "",  // 机器人账号id
        "type": 1,  // 1/文本消息 3/图片消息 34/语音消息  42/名片消息  43/视频 47/动态表情 48/地理位置  49/分享链接  2001/红包  2002/小程序  2003/群邀请  更多请参考常量表
        "from_wxid": "",  // 来源用户ID
        "from_name": "",  // 来源用户昵称
 		"from_group": "",  // 来源群号
        "from_group_name": "",  // 来源群名称
        "msg": "",  // 消息内容
        "msgid": 0  // 消息ID
}
*/
type GroupChatResp struct {
	Appkey      string      `json:"appkey"`
	Event       string      `json:"event"`
	Robot       string      `json:"robot_wxid"`
	Type        string      `json:"type"`
	From        string      `json:"from_wxid"`
	NickName    string      `json:"from_name"`
	Content     interface{} `json:"msg"`
	Useringroup string      `json:"useringroup"`
	Group       string      `json:"from_group"`
	GroupName   string      `json:"from_group_name"`
	MsgId       string      `json:"msgid"`
}
