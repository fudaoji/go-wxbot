package install

import (
	. "go-wxbot/db"
	. "go-wxbot/model"
)

//  初始化数据库表
func InstallHandle() {
	MysqlClient.AutoMigrate(&Appkey{})
}
