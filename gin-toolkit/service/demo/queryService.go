package demo

import (
	"github.com/gin-gonic/gin"
)

func Call(empName string, cityId int, c *gin.Context) (empNameRes string, cityIdRes int, err error) {
	empNameRes = "empNameRes"
	cityIdRes = 2333
	return empNameRes, cityIdRes, nil
}