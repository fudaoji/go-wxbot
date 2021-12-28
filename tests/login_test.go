package tests

import (
	"encoding/json"
	"fmt"
	"go-wxbot/core"
	"testing"
)

// 登录请求体
type loginRes struct {
	// 回调地址
	Webhook string `form:"webhook" json:"webhook"`
}

func TestGetLoginCode(t *testing.T) {
	method, url := "GET", Apis["getlogincode"]
	res := Request(method, url, nil)
	fmt.Printf("%#v", res)
}

func TestCheckLogin(t *testing.T) {
	method, url := "POST", Apis["checklogin"]
	res := Request(method, url, &loginRes{Webhook: "http://wxbot.oudewa.cn/admin/onmessage/botCallback"})
	fmt.Printf("%#v", res)
}

func TestJson(t *testing.T) {
	params := map[string]interface{}{"name": "doogie", "age": 12}
	fmt.Println(len(params))
	res, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
	}
	t.Logf("%#v", string(res))
	fmt.Printf("%#v", string(res))
}

// 登录请求体
type lRes struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Verify   string `form:"verify" json:"verify"`
}

func TestPost(t *testing.T) {
	res := core.ReqPostJson("http://wxbot.oudewa.cn/admin/auth/login", &lRes{Username: "admin", Password: "123456", Verify: "123456"}, nil)

	fmt.Printf("%#v", res)
}
