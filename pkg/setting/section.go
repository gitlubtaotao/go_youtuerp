package setting

import "time"

type AppServiceSettings struct {
}

type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type ServerSettingS struct {
	ReadTimeout time.Duration
}

type QiniuUploadSettingS struct {
	QiniuAccessKey  string `json:"qiniu_access_key" yaml:"qiniu_access_key"`
	QiniuSecretKey  string `json:"qiniu_secret_key" yaml:"qiniu_secret_key"`
	QiniuFileServer string `json:"qiniu_file_server" yaml:"qiniu_file_server"`
	QiniuBucket     string `yaml:"qiniu_bucket" json:"qiniu_bucket"`
}
