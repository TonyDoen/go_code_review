package demo

import (
	"../../common"
	"../../service/demo"
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

	empNameRes, cityIdRes, err := demo.Call("tony", 666, c)
	if err != nil {
		common.RenderJsonFail(c, err)
		return
	}

	var result = make(map[string]interface{})
	result["empNameRes"] = empNameRes
	result["cityIdRes"] = cityIdRes
	common.RenderJsonSuccess(c, gin.H{
		"result":result,
	})
	return
}