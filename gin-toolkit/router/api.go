package router

import (
	"../controller/demo"
	"github.com/gin-gonic/gin"
)

func Api(r *gin.Engine) {
	demoGroup := r.Group("/api/demo")
	demoGroup.GET("/query", demo.Query)
}
