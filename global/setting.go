package global

import (
	"github.com/go-redis/redis/v7"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
	"time"
	"youtuerp/pkg/setting"
	redis2 "youtuerp/redis"
)

var (
	DataEngine         *gorm.DB                     // database engine
	RedisEngine        *redis.Client                // redis engine
	IrisAppEngine      *iris.Application            // iris app engine
	RedSetting         = redis2.NewRedis()          // redis setting instance
	AppServiceSetting  *setting.AppServiceSettings  // app service setting
	ServerSetting      *setting.ServerSettingS      // server setting
	DatabaseSetting    *setting.DatabaseSettingS    // database setting
	QiNiuUploadSetting *setting.QiniuUploadSettingS // qiniu upload setting
)

func SetupCommonSetting() error {
	set, err := setting.NewSetting()
	if err != nil {
		return err
	}
	if err = set.ReadSection("Server", &ServerSetting); err != nil {
		return err
	}
	if err = set.ReadSection("App", &AppServiceSetting); err != nil {
		return err
	}
	if err = set.ReadSection("Database", &DatabaseSetting); err != nil {
		return err
	}
	if err = set.ReadSection("Email", &QiNiuUploadSetting); err != nil {
		return err
	}
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second
	return nil
}
