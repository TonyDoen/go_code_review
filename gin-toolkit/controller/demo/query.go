package demo

import (
	"github.com/TonyDoen/go_code_review/gin-toolkit/common"
	"github.com/TonyDoen/go_code_review/gin-toolkit/service/demo"
	"github.com/gin-gonic/gin"
)

type paramQueryList struct {
	EmpName string `form:"EmpName" binding:"required,max=32,min=1" `
	CityId  string `form:"CityId" binding:"required,max=32,min=1" `
}

func Query(c *gin.Context) {
	var params paramQueryList
	if err := c.ShouldBind(&params); err != nil {
		common.RenderJsonFail(c, err)
		return
	}

	empNameRes, cityIdRes, err := demo.Call(c, "tony", 666)
	if err != nil {
		common.RenderJsonFail(c, err)
		return
	}

	var result = make(map[string]interface{})
	result["empNameRes"] = empNameRes
	result["cityIdRes"] = cityIdRes
	common.RenderJsonSuccess(c, gin.H{
		"result": result,
	})
	return
}

func QueryAll(c *gin.Context) {
	res, err := demo.GetAll(c)
	if nil != err {
		common.RenderJsonFail(c, err)
	}
	common.RenderJsonSuccess(c, gin.H{"list": res})
}
