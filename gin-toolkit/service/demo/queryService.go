package demo

import (
	"encoding/json"
	"fmt"
	"github.com/TonyDoen/go_code_review/gin-toolkit/common"
	"github.com/TonyDoen/go_code_review/gin-toolkit/conf/data"
	"github.com/TonyDoen/go_code_review/gin-toolkit/model"
	"github.com/gin-gonic/gin"
	"time"
)

const duration = 15 * time.Minute

func Call(c *gin.Context, empName string, cityId int) (empNameRes string, cityIdRes int, err error) {

	var one = model.Demo{Name: "xiaoming", Desc: "2333", CreateTime: time.Now()}
	err = one.Insert()
	if nil != err {
		return empNameRes, cityIdRes, err
	}

	var oneName = one.Name
	// delete in redis
	_, err = data.RedisClient.Del(oneName).Result()
	if nil != err {
		return empNameRes, cityIdRes, err
	}

	// add in redis
	info, _ := json.Marshal(one)
	data.RedisClient.Restore(oneName, duration, string(info))

	// query in mysql
	thisOne, err := model.GetDemoByName(oneName)
	if nil != err {
		return empNameRes, cityIdRes, err
	}

	common.DebugLogger(c, thisOne)

	empNameRes = thisOne.Name + empName
	cityIdRes = 2333 + cityId

	return empNameRes, cityIdRes, nil
}

func GetAll(c *gin.Context) (res []model.Demo, err error) {
	res, err = model.GetDemoAllList()
	if nil != err {
		return res, err
	}

	for i, one := range res {
		if 1 == i {
			err = model.UpdateDescByName(fmt.Sprintf("第%d名", i), one.Name)
		} else {
			one.Desc = fmt.Sprintf("第%d名", i)
			one.Name = fmt.Sprintf("第%d名的%s", i, one.Name)
			err = one.Update()
		}
	}
	res, err = model.GetDemoAllList()
	return res, err
}

func CronCall() {
	res, err := model.GetDemoAllList()
	if nil != err {
		return
	}

	for _, one := range res {
		_, err = data.RedisClient.Del(one.Name).Result()
	}
}
