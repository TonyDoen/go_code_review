package main

import (
	"./common"
	"./conf"
	"./conf/data"
	"./router"
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
