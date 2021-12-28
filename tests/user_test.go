package tests

import (
	"fmt"
	"testing"
)

// 修改备注名请求体
type setRemarkNameRes struct {
	// 用户名
	To string `form:"to" json:"to"`
	// 正文
	RemarkName string `form:"remark_name" json:"remark_name"`
}

func TestSetFriendRemarkName(t *testing.T) {
	method, url := "POST", Apis["setfriendremarkname"]
	res := Request(method, url, &setRemarkNameRes{To: "@59d722cbf25d962039d0239dec951c96f7df2ad79a196ed4b561140d6161c3a8", RemarkName: "傅道集"})
	fmt.Printf("%#v", res)
}

func TestListGroups(t *testing.T) {
	method, url := "GET", Apis["listgroups"]
	res := Request(method, url, nil)
	fmt.Printf("%#v", res)
}

func TestListFriends(t *testing.T) {
	method, url := "GET", Apis["listfriends"]
	res := Request(method, url, nil)
	fmt.Printf("%#v", res)
}

func TestGetCurrentUser(t *testing.T) {
	method, url := "GET", Apis["getcurrentuser"]
	res := Request(method, url, nil)
	fmt.Printf("%#v", res)
}
