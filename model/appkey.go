package model

import (
	. "go-wxbot/db"
)

const tableName = "appkey"

// Appkey appkey表
type Appkey struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement:true;comment:'ID'"`
	Appkey   string `json:"appkey" gorm:"type:varchar(32);comment:'appKey'"`
	Deadline int64  `json:"deadline" gorm:"type:int(10);comment:'到期时间'"`
	Webhook  string `json:"webhook" gorm:"type:varchar(200);comment:'事件回调地址'"`
	Uin      int64  `json:"uin" gorm:"type:bigint(20);comment:'uin'"`
}

//findByUin 根据uin找记录
func (a *Appkey) FindByUin() {
	MysqlClient.Table(tableName).Where("uin = ?", a.Uin).First(a)
}

//findByAppkey 根据appkey找记录
func (a *Appkey) FindByAppkey() {
	MysqlClient.Table(tableName).Where("appkey = ?", a.Appkey).First(a)
}

//UpdateWebhookByAppkey 根据appkey设置webhook
func UpdateWebhookByAppkey(appkey string, webhook string, uin int64) {
	MysqlClient.Model(&Appkey{}).Where("appkey = ?", appkey).
		Updates(Appkey{Webhook: webhook, Uin: uin})
}
