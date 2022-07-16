package core

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

// RedisConfig Redis配置
var (
	RedisConfig redisConfig
	MySQLConfig mysqlConfig
)

// Redis配置
type redisConfig struct {
	Host     string // Redis主机
	Port     string // Redis端口
	Password string // Redis密码
	Db       int    // Redis库
}

// MySQL配置
type mysqlConfig struct {
	Host     string // 主机
	Port     string // 端口
	Username string // 用户名
	Password string // 密码
	DbName   string // 数据库名称
}

type SocketConfig struct {
	Host   string // 主机
	Port   string // 端口
	Scheme string // scheme
	Path   string //
}

// InitRedisConfig 初始化Redis配置
func InitMysqlConfig() {
	// Mysql配置
	//主机
	host := GetVal("db.host", "127.0.0.1")
	// 端口
	port := GetVal("db.port", "3306")
	// 密码
	password := GetVal("db.password", "")
	// 数据库
	database := GetVal("db.database", "test")
	//用户名
	username := GetVal("db.username", "root")

	MySQLConfig = mysqlConfig{
		Host:     host,
		Port:     port,
		Password: password,
		Username: username,
		DbName:   database,
	}
}

// InitRedisConfig 初始化Redis配置
func InitRedisConfig() {
	// RedisHost Redis主机
	//host := utils.GetEnvVal("REDIS_HOST", "127.0.0.1")
	host := GetVal("redis.host", "127.0.0.1")
	// RedisPort Redis端口
	//port := utils.GetEnvVal("REDIS_PORT", "6379")
	port := GetVal("redis.port", "6379")
	// RedisPassword Redis密码
	//password := utils.GetEnvVal("REDIS_PWD", "")
	password := GetVal("redis.pwd", "")
	// Redis库
	//db := utils.GetEnvIntVal("REDIS_DB", 0)
	db := GetIntVal("redis.db", 0)

	RedisConfig = redisConfig{
		Host:     host,
		Port:     port,
		Password: password,
		Db:       db,
	}
}

//InitConfig 读取配置文件
func InitConfig() {
	mode := flag.String("mode", "dev", "dev mode")
	flag.Parse()
	viper.SetConfigFile(fmt.Sprintf("./setting_%s.yaml", *mode))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

//GetVal 获取配置文件的字符串配置值
func GetVal(key string, defaultVal string) string {
	if viper.IsSet(key) {
		return viper.GetString(key)
	}
	return defaultVal
}

//GetIntVal 获取配置文件的整型配置值
func GetIntVal(key string, defaultVal int) int {
	if viper.IsSet(key) {
		return viper.GetInt(key)
	}
	return defaultVal
}
