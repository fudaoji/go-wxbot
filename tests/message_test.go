package tests

import (
	"fmt"
	"testing"
)

// 发送消息请求体
type sendMsg struct {
	// 送达人UserName
	To string `form:"to" json:"to"`
	// 正文
	Content string `form:"msg" json:"msg"`
}
type sendFileMsg struct {
	// 送达人UserName
	To string `form:"to" json:"to"`
	// 正文
	Content string `form:"path" json:"path"`
}

const (
	dj   string = "@a3eb65322a011f97a312a13ac8db53997d47079c4edac5c2c0cd5ea8729036a1"
	csq  string = "@@bf763a43269b85e3a6525df1cba4d6a2170dafccaa5c5bd8a40242cda81d5c42"
	jane string = "@ec282e3c4f37df78fedf8f8e2d57287dcac256b2bde01e1054887cd95eaf3656"
	q53  string = "@@0d77e58642ee94b43a1440977759e5419aab2c3d21f56effd5398b1ef5754142"
)

//发送视频
func TestVideo(t *testing.T) {
	method, url := "POST", Apis["sendVideo"]
	video := "https://zmzgz.images.huihuiba.net/20191113/1/1_uxTZBI17_Wildlife.mp4"
	res := Request(method, url, sendFileMsg{To: dj, Content: video})
	fmt.Printf("%#v", res)
}

//发送文件
func TestFile(t *testing.T) {
	method, url := "POST", Apis["sendFile"]
	file := "https://zmzgz.images.huihuiba.net/20191113/1/1_uxTZBI17_Wildlife.mp4"
	username := "@3e70e3ac94e107ddfbe97fa57fa26d7b7b08a9a982f2abda8ddda372ba0b56b6"
	res := Request(method, url, sendFileMsg{To: username, Content: file})
	fmt.Printf("%#v", res)
}

//发送图片
func TestImg(t *testing.T) {
	method, url := "POST", Apis["sendImg"]
	img := "https://zyx.images.huihuiba.net/0-5b84f6adbded5.png"
	username := q53
	res := Request(method, url, sendMsg{To: username, Content: img})
	fmt.Printf("%#v", res)
}

//发送文本消息
func TestText(t *testing.T) {
	method, url := "POST", Apis["sendText"]
	res := Request(method, url, sendMsg{To: csq, Content: "hi"})
	fmt.Printf("%#v", res)
}
