package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"youtuerp/pkg/setting"
)

//设置mysql数据的连接
func MysqlDBDialect(setting *setting.DatabaseSettingS) gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		setting.Username,
		setting.Password,
		setting.Host,
		setting.DBName,
		setting.Charset,
		setting.ParseTime,
	)
	return mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	})
}
