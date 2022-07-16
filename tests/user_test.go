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

// 邀请好友入群
type inviteIngroupRes struct {
	// 用户名
	Group string `form:"group" json:"group"`
	// 正文
	Friends []string `form:"friends" json:"friends"`
}

// 邀请好友入群
type removeMembers struct {
	// 用户名
	Group string `form:"group" json:"group"`
	// 正文
	Members []string `form:"members" json:"members"`
}

func TestRemoveGroupMembers(t *testing.T) {
	method, url := "POST", Apis["removegroupmembers"]
	members := [...]string{dj}
	res := Request(method, url, &removeMembers{Group: q53, Members: members[:]})
	fmt.Printf("%#v", res)
}

func TestGetGroupMembers(t *testing.T) {
	method, url := "POST", Apis["getgroupmembers"]
	res := Request(method, url, &inviteIngroupRes{Group: q53})
	fmt.Printf("%#v", res)
}

func TestInviteInGroup(t *testing.T) {
	method, url := "POST", Apis["inviteingroup"]
	friends := [...]string{jane}
	res := Request(method, url, &inviteIngroupRes{Group: csq, Friends: friends[:]})
	fmt.Printf("%#v", res)
}

func TestSetFriendRemarkName(t *testing.T) {
	method, url := "POST", Apis["setfriendremarkname"]
	res := Request(method, url, &setRemarkNameRes{To: csq, RemarkName: "采品测试器1"})
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
