package tests

import (
	"fmt"
	"testing"
)

// 发送消息请求体
type sendMsg struct {
	// 送达人UserName
	To string
	// 正文
	Content string
}

//发送视频
func TestVideo(t *testing.T) {
	method, url := "POST", Apis["sendVideo"]
	video := "https://zmzgz.images.huihuiba.net/20191113/1/1_uxTZBI17_Wildlife.mp4"
	username := "@3e70e3ac94e107ddfbe97fa57fa26d7b7b08a9a982f2abda8ddda372ba0b56b6"
	res := Request(method, url, sendMsg{To: username, Content: video})
	fmt.Printf("%#v", res)
}

//发送文件
func TestFile(t *testing.T) {
	method, url := "POST", Apis["sendFile"]
	file := "https://zmzgz.images.huihuiba.net/20191113/1/1_uxTZBI17_Wildlife.mp4"
	username := "@3e70e3ac94e107ddfbe97fa57fa26d7b7b08a9a982f2abda8ddda372ba0b56b6"
	res := Request(method, url, sendMsg{To: username, Content: file})
	fmt.Printf("%#v", res)
}

//发送图片
func TestImg(t *testing.T) {
	method, url := "POST", Apis["sendImg"]
	img := "https://zyx.images.huihuiba.net/0-5b84f6adbded5.png"
	username := "@3e70e3ac94e107ddfbe97fa57fa26d7b7b08a9a982f2abda8ddda372ba0b56b6"
	res := Request(method, url, sendMsg{To: username, Content: img})
	fmt.Printf("%#v", res)
}

//发送文本消息
func TestText(t *testing.T) {
	method, url := "POST", Apis["sendText"]
	res := Request(method, url, sendMsg{To: "@@5dda033a3c603d7b540728382289ec631e61eee29cd44e28470a12b5883d98c9", Content: "很好"})
	fmt.Printf("%#v", res)
}
