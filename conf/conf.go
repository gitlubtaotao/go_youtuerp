package conf

import (
	"encoding/json"
	"github.com/go-redis/redis/v7"
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
	ReisCon       *redis.Client
)

type EnvironmentConfig struct {
	Env             string `json:"env" yaml:"env"`
	PerPage         uint   `json:"per_page" yaml:"per_page"`
	DSN             string `json:"dsn" yaml:"dsn"`
	TokenSecret     string `json:"token_secret" yaml:"token_secret"`
	AllowedOrigins  string `json:"allowed_origins" yaml:"allowed_origins"`
	ExpireTime      int    `json:"expire_time" yaml:"expire_time"`
	QiniuAccessKey  string `json:"qiniu_access_key" yaml:"qiniu_access_key"`
	QiniuSecretKey  string `json:"qiniu_secret_key" yaml:"qiniu_secret_key"`
	QiniuFileServer string `json:"qiniu_file_server" yaml:"qiniu_file_server"`
	QiniuBucket     string `yaml:"qiniu_bucket" json:"qiniu_bucket"`
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


