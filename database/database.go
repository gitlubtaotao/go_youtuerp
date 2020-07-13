package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"youtuerp/conf"
)

var dataEngine *gorm.DB
func GetDBCon() *gorm.DB {
	return dataEngine
}
type DataBase struct {

}

/*
 * 初始化项目
 */
func (d *DataBase) DefaultInit() error {
	var err error
	if err = d.InitDataBase(); err != nil {
		return err
	}
	if err = d.Migration(); err != nil {
		return err
	}
	return nil
}

func (d *DataBase) InitDataBase() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,         // Disable color
		},
	)
	var err error
	dataEngine, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               conf.Configuration.DSN,
		DefaultStringSize: 256,
	}), &gorm.Config{
		Logger: newLogger,
		PrepareStmt: true,
	})
	if err != nil {
		return err
	}
	sqlDB, err := dataEngine.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(1200)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err = d.Migration(); err != nil {
		return err
	}
	return nil
}

/*
 * 注册迁移文件
 */


func (d *DataBase) Migration() error {
	return nil
}
