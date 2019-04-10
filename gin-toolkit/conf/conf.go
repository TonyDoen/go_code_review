package conf

import (
	"io/ioutil"
	"path"
	"runtime"

	"github.com/TonyDoen/go_code_review/gin-toolkit/common"
	"gopkg.in/yaml.v2"
)

//支持数组，工具
type TConf struct {
	Redis struct {
		Addr string `yaml:"addr"`
	}
	Mysql struct {
		Addr         string `yaml:"addr"`
		User         string `yaml:"user"`
		PassWord     string `yaml:"password"`
		DataBase     string `yaml:"database"`
		MaxIdleConns int    `yaml:"maxidleconns"`
		MaxOpenConns int    `yaml:"maxopenconns"`
	}
	Port string
	Nsq  struct {
		Addr  string `yaml:"addr"`
		Topic string `yaml:"topic"`
	}
}

var Conf TConf

func init() {
	yamlFile, err := ioutil.ReadFile(getCurrentPath() + "/conf.yaml")
	if err != nil {
		common.Panic4Logger(nil, "yamlfile get error: %v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	if err != nil {
		common.Panic4Logger(nil, "yaml unmarshal error: %v", err)
	}
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

