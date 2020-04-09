package initialize

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
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
	if err := d.InitDataBase(); err != nil {
		return err
	}
	d.Migration()
	return nil
}
func (d *DataBase) InitDataBase() error {
	var err error
	DataEngine, err = gorm.Open("mysql", "root:qweqwe123@tcp(127.0.0.1:3306)/go_youtuerp?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true")
	if err != nil {
		return err
	}
	defer DataEngine.Close()
	DataEngine.DB().SetMaxOpenConns(1200)
	DataEngine.LogMode(true)
	DataEngine.SetLogger(log.New(os.Stdout, "\r\n", 0))
	d.Migration()
	return nil
}

/*
 * 注册迁移文件
 */
func (d *DataBase) Migration() {

}
