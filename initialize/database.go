package initialize

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"youtuerp/conf"
)

var DataEngine *gorm.DB

type IDataBase interface {
	DefaultInit()
	Migration()
	InitDataBase()
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
	var err error
	DataEngine, err = gorm.Open("mysql", conf.Configuration.DSN)
	if err != nil {
		return err
	}
	defer DataEngine.Close()
	DataEngine.DB().SetMaxOpenConns(1200)
	DataEngine.LogMode(true)
	DataEngine.SetLogger(log.New(os.Stdout, "\r\n", 0))
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
