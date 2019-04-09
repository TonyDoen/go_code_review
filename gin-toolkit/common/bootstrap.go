package common

import (
	"github.com/gin-gonic/gin"
)

func Bootstrap(router *gin.Engine) {
	//环境判断 env GIN_MODE=release/debug
	if GetRunEnv() == EnvProduct {
		gin.SetMode(gin.ReleaseMode)
	}
	//日志配置json格式，默认标准输出，初始化日志
	InitLog()

	//加载异常处理
	router.Use(Recovery())

	//日志中间件
	router.Use(LoggerTemplate())

	//health
	router.GET("/health", func(c *gin.Context) {
		c.String(200, "success")
	})
	router.HEAD("/health", func(c *gin.Context) {
		c.String(200, "success")
	})

	//性能分析工具
	if IsPprofEnv() {
		PprofRegister(router)
	}
}
