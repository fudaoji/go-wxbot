## WebWechat

基于`Golang`语言和`Gin`框架的个人微信机器人，以RESTful Api方式对外提供服务。

项目初期主要参考<a href="https://github.com/lixh00/web-wechat" target="_blank">Web-WeChat</a>，后面根据自己的需求进行设计、完善。

## 安装

```shell
# 下载代码
git clone https://github.com/fudaoji/go-wxbot.git
# 更新依赖
go mod download
# 复制一份生产环境配置文件修改配置文件
cp setting_dev.yaml  setting_prod.yaml
# 编译
go build main.go
# 清理无用mod引用
go mod tidy
```
## 运行

```shell
# 运行测试环境
./main
# 运行生成环境
./main  -mode="prod"
```

更多使用说明请访问[开发文档](https://kyphp.kuryun.com/home/guide/bot.html)

## Thanks

<a href="https://github.com/lixh00/web-wechat" target="_blank">Web-WeChat</a>

<a href="https://github.com/eatmoreapple/openwechat" target="_blank">OpenWeChat</a>
