package conf

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

/*
 * 系统的配置文件
 * 设置系统默认的配置文件
 */

var (
	Configuration *EnvironmentConfig
	IrisApp       *iris.Application
)

type EnvironmentConfig struct {
	PerPage uint   `json:"per_page" yaml:"per_page"`
	DSN     string `json:"dsn" yaml:"dsn"`
}

/*
 * 初始化系统环境变量设置
 */
func NewSysConfig(env string) (err error) {
	fileReader, err := ioutil.ReadFile("../conf/conf.yaml")
	if err != nil {
		return err
	}
	attr := make(map[string]interface{})
	err = yaml.Unmarshal(fileReader, &attr)
	if err != nil {
		return err
	}
	data := make(map[string]interface{})
	for k, v := range attr[env].(map[interface{}]interface{}) {
		s := k.(string)
		data[s] = v
	}
	bytes, err := json.Marshal(data)
	err = json.Unmarshal(bytes, &Configuration)
	return err
}
