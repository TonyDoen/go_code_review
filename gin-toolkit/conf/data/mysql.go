package data

import (
	"fmt"
	"github.com/TonyDoen/go_code_review/gin-toolkit/common"
	"github.com/TonyDoen/go_code_review/gin-toolkit/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"time"
)

var MysqlClient *gorm.DB

func InitMysql() {
	var err error
	MysqlClient, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai",
		conf.Conf.Mysql.User,
		conf.Conf.Mysql.PassWord,
		conf.Conf.Mysql.Addr,
		conf.Conf.Mysql.DataBase))

	if err != nil {
		common.Panic4Logger(nil, "mysql connect error: %v", err)
	}

	MysqlClient.DB().SetMaxIdleConns(conf.Conf.Mysql.MaxIdleConns)
	MysqlClient.DB().SetMaxOpenConns(conf.Conf.Mysql.MaxOpenConns)

	MysqlClient.DB().SetConnMaxLifetime(3600 * time.Second)
	MysqlClient.SetLogger(&gormLogger{})
	MysqlClient.LogMode(true)
}

type gormLogger struct{}

func (*gormLogger) Print(v ...interface{}) {
	logrus.WithFields(logrus.Fields{"module": "gorm"}).Debug(v...)
}
