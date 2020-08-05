package setting

import "time"

type AppSettingS struct {
	TokenSecret    string
	PerPage        string
	AllowedOrigins []string
	ExpireTime     time.Duration
	Env            string
}
type DatabaseSettingS struct {
	DBType       string
	Username     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	Collation    string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type ServerSettingS struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	RunMode      string
	HttpPort     string
}

type QiniuUploadSettingS struct {
	QiniuAccessKey  string `json:"qiniu_access_key" yaml:"qiniu_access_key"`
	QiniuSecretKey  string `json:"qiniu_secret_key" yaml:"qiniu_secret_key"`
	QiniuFileServer string `json:"qiniu_file_server" yaml:"qiniu_file_server"`
	QiniuBucket     string `yaml:"qiniu_bucket" json:"qiniu_bucket"`
}
