package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"youtuerp/models"
)

var dataEngine *gorm.DB

func GetDBCon() *gorm.DB {
	return dataEngine
}

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
	password := os.Getenv("MYSQL_ROOT_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	fmt.Print(password, database, "ssss")
	dsn := "root:qweqwe123@tcp(db:3306)/go_youtuerp?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=true"
	dataEngine, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	dataEngine.DB().SetMaxOpenConns(1200)
	dataEngine.LogMode(true)
	dataEngine.SetLogger(log.New(os.Stdout, "\r\n", 0))
	if err = d.Migration(); err != nil {
		return err
	}
	return nil
}

/*
 * 注册迁移文件
 */
func (d *DataBase) Migration() error {
	if GetDBCon().HasTable("companies") {
		GetDBCon().AutoMigrate(&models.CompanyInfo{})
	}
	GetDBCon().AutoMigrate(&models.Employee{}, &models.UserCompany{},
		&models.Department{}, &models.CrmCompany{}, &models.Contact{}, &models.Department{})
	return nil
}
