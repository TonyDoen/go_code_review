package main

import (
	"github.com/TonyDoen/go_code_review/gin-toolkit/common"
	"github.com/TonyDoen/go_code_review/gin-toolkit/conf"
	"github.com/TonyDoen/go_code_review/gin-toolkit/conf/data"
	"github.com/TonyDoen/go_code_review/gin-toolkit/router"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin
	r := gin.New()

	// 启动配置
	common.Bootstrap(r)

	data.InitMysql()
	data.InitRedis()

	// 路由
	router.Api(r)

	// 定时任务
	router.Crontab()

	_ = endless.ListenAndServe(":"+conf.Conf.Port, r)
}
