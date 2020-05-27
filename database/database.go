package database

import (
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"youtuerp/conf"
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
	dataEngine, err = gorm.Open("mysql", conf.Configuration.DSN)
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
	db := GetDBCon()
	db.AutoMigrate(&models.Company{}, &models.User{},
		&models.Account{}, &models.Role{}, &models.Invoice{}, &models.Address{})
	db.AutoMigrate(&models.Department{})
	db.AutoMigrate(&models.CrmClue{}, &models.CrmTrack{})
	db.AutoMigrate(&models.Setting{}, &models.NumberSetting{})
	db.AutoMigrate(&models.FinanceFeeType{}, &models.FinanceRate{})
	db.AutoMigrate(&models.BaseDataLevel{}, &models.BaseDataCode{},
		&models.BaseDataPort{}, &models.BaseDataCarrier{}, &models.BaseWarehouse{})
	db.AutoMigrate(
		&models.OrderMaster{},
		&models.OrderExtendInfo{},
		&models.FormerSeaInstruction{},&models.FormerSeaBook{},
		&models.FormerSeaSoNo{},&models.SeaCargoInfo{},&models.SeaCapList{},
	)
	return nil
}
