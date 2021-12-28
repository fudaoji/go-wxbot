package middleware

import (
	"go-wxbot/core"
	"go-wxbot/global"
	"strings"
	"time"

	"go-wxbot/model"

	"github.com/gin-gonic/gin"
)

// CheckAppKeyIsLoggedInMiddleware 检查AppKey是否已登录微信
func CheckAppKeyIsLoggedInMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appKey := ctx.Request.Header.Get("AppKey")

		flag := true
		// 判断AppKey是否存在，商业化后根据appkey收费
		if !checkAppKey(appKey) {
			core.FailWithMessage("AppKey非法或已过期", ctx)
			flag = false
		} else if !strings.Contains(ctx.Request.RequestURI, "login") {
			// 如果不是登录请求，判断AppKey是否有效
			if err := global.CheckBot(appKey); err != nil {
				core.FailWithMessage("AppKey预检失败："+err.Error(), ctx)
				flag = false
			}
		}
		if flag {
			ctx.Next()
		} else {
			ctx.Abort()
		}
	}
}

// CheckAppKeyExistMiddleware 检查是否有appKey
func CheckAppKeyExistMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		appKey := ctx.Request.Header.Get("AppKey")
		// 判断AppKey是否存在，商业化后根据appkey收费
		if !checkAppKey(appKey) {
			core.FailWithMessage("AppKey非法或已过期", ctx)
			ctx.Abort()
		} else if len(appKey) < 1 { // 先判断AppKey是不是传了
			core.FailWithMessage("AppKey为必传参数", ctx)
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}

//验证AppKey是否存在且未过期
func checkAppKey(appKey string) bool {
	//exists, _ := utils.ContainsStr(global.AppKeys, appKey)
	//return exists
	var key = model.Appkey{Appkey: appKey}
	key.FindByAppkey()
	return key.Deadline > time.Now().Unix()
}
