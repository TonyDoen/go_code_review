package router

import (
	"github.com/TonyDoen/go_code_review/gin-toolkit/controller/demo"
	"github.com/gin-gonic/gin"
)

func Api(r *gin.Engine) {
	demoGroup := r.Group("/api/demo")
	demoGroup.GET("/query", demo.Query)
	demoGroup.GET("/queryAll", demo.QueryAll)
}
