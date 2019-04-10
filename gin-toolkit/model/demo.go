package model

import (
	"github.com/TonyDoen/go_code_review/gin-toolkit/common"
	"github.com/TonyDoen/go_code_review/gin-toolkit/conf/data"
	"time"
)

type Demo struct {
	ID         int       `gorm:"primary_key;column:id" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Desc       string    `gorm:"column:desc" json:"desc"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
}

func (Demo) TableName() string {
	return "demo"
}

func (p *Demo) Insert() (err error) {
	err = data.MysqlClient.Create(p).Error
	if nil != err {
		common.SetErrPrint4Msg(&common.DbDemoInsertError, err.Error())
	}
	return err
}

func (p *Demo) Update() (err error) {
	err = data.MysqlClient.Model(&Demo{}).Where("id = ?", p.ID).Updates(map[string]interface{}{
		"name": p.Name,
		"desc": p.Desc,
		"create_time": time.Now(),
	}).Error
	if nil != err {
		common.SetErrPrint4Msg(&common.DbDemoUpdateError, err.Error())
	}
	return err
}

func GetDemoByName(name string) (res Demo, err error) {
	err = data.MysqlClient.Where("name = ?", name).Find(&res).Error
	if nil != err {
		common.SetErrPrint4Msg(&common.DbDemoGetError, err.Error())
	}
	return res, err
}

func GetDemoAllList() (res []Demo, err error) {
	err = data.MysqlClient.Find(&res).Error
	if nil != err {
		common.SetErrPrint4Msg(&common.DbDemoGetError, err.Error())
	}
	return res, err
}

func UpdateDescByName(desc, name string) (err error) {
	err = data.MysqlClient.Model(&Demo{}).Where("name = ?", name).Updates(map[string]interface{}{
		"desc": desc,
	}).Error
	if nil != err {
		common.SetErrPrint4Msg(&common.DbDemoUpdateError, err.Error())
	}
	return err
}
