package conf

import (
	"../common"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type App struct {
	AppKey  string `yaml:"appkey"`
	Domain  string `yaml:"domain"`
	Timeout string `yaml:"timeout"`
}

type TApi struct {
	Demo App
}

var Api TApi

func init() {
	yamlFile, err := ioutil.ReadFile(getCurrentPath() + "/api.yaml")
	if err != nil {
		common.Panic4Logger(nil, "yaml file get error: %v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &Api)
	if err != nil {
		common.Panic4Logger(nil, "yaml unmarshal error: %v", err)
	}
}
