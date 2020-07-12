package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
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
	dataEngine, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               conf.Configuration.DSN,
		DefaultStringSize: 256,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
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
	db := GetDBCon()
	db.AutoMigrate(&models.Company{}, &models.User{},
		&models.Account{}, &models.Role{}, &models.Invoice{}, &models.Address{})
	db.AutoMigrate(&models.Department{})
	db.AutoMigrate(&models.CrmClue{}, &models.CrmTrack{})
	db.AutoMigrate(&models.Setting{}, &models.NumberSetting{},
		&models.NumberSettingHistory{})
	db.AutoMigrate(&models.FinanceFeeType{}, &models.FinanceRate{})
	db.AutoMigrate(&models.BaseDataLevel{}, &models.BaseDataCode{},
		&models.BaseDataPort{}, &models.BaseDataCarrier{}, &models.BaseWarehouse{})
	db.AutoMigrate(
		&models.OrderMaster{},
		&models.OrderExtendInfo{},
		&models.FormerSeaInstruction{}, &models.FormerSeaBook{},
		&models.FormerSeaSoNo{}, &models.SeaCargoInfo{}, &models.SeaCapList{},
		&models.FormerOtherService{}, &models.FormerTrailerOrder{}, &models.TrailerCabinetNumber{},
		&models.FormerWarehouseService{}, &models.FormerCustomClearance{},
	)
	db.AutoMigrate(&models.FinanceFee{})
	db.AutoMigrate(&models.Attachment{})
	return nil
}
