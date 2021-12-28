package tests

import (
	"fmt"
	"testing"
)

// 发送消息请求体
type sendMsg struct {
	// 送达人UserName
	To string
	// 消息类型
	Type string
	// 正文
	Content string
}

//发送图片给好友
func TestImgToFriend(t *testing.T) {
	method, url := "POST", Apis["imgToFriend"]
	img := "https://zyx.images.huihuiba.net/0-5b84f6adbded5.png"
	username := "@67822500676b2e82af69301bfa7e1c3df79f05e7c686238e6e47966d61a53fa6"
	res := Request(method, url, sendMsg{To: username, Type: "image", Content: img})
	fmt.Printf("%#v", res)
}

//发送消息给好友
func TestTextToFriend(t *testing.T) {
	method, url := "POST", Apis["msgToFriend"]
	res := Request(method, url, sendMsg{To: "@86b5c76331d9a825c68ed9aff439cf9114bcaae01e0603c7700b1b59645ed290", Type: "text", Content: "test"})
	fmt.Printf("%#v", res)
}

//发送消息给群聊
func TestTextToGroup(t *testing.T) {
	method, url := "POST", Apis["msgToGroup"]
	res := Request(method, url, sendMsg{To: "有家、有爱", Type: "text", Content: "hi"})
	fmt.Printf("%#v", res)
}
