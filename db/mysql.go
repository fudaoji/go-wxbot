package db

import (
	"fmt"
	"go-wxbot/core"
	"go-wxbot/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type mysqlDB struct {
	*gorm.DB
}

var MysqlClient mysqlDB

// InitMongoConnHandle 初始化MongoDB连接
func InitMysqlConnHandle() {
	// 读取配置
	core.InitMysqlConfig()

	/* db, err := gorm.Open("mysql",
	fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&loc=Local",
		core.MySQLConfig.Username, core.MySQLConfig.Password, core.MySQLConfig.Host, core.MySQLConfig.Port, core.MySQLConfig.
			DbName)) */
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		core.MySQLConfig.Username, core.MySQLConfig.Password, core.MySQLConfig.Host, core.MySQLConfig.Port, core.MySQLConfig.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic("failed to connect mysql")
	}

	logger.Log.Info("MysqlDB连接初始化成功")
	MysqlClient = mysqlDB{db}
}
