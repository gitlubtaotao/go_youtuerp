package global

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
	"youtuerp/pkg/database"
)

//global database connection
func GetDBCon() *gorm.DB {
	return DataEngine
}

// mysql database connection engine
func NewDBEngine() error {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // Disable color
		},
	)
	var (
		dialect gorm.Dialector
		err     error
	)
	if DatabaseSetting.DBType == "mysql" {
		dialect = database.MysqlDBDialect(DatabaseSetting)
	}
	DataEngine, err = gorm.Open(dialect, &gorm.Config{
		PrepareStmt: true,
		Logger:      newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   DatabaseSetting.TablePrefix,
			SingularTable: false,
		},
	})
	if err != nil {
		return err
	}
	if sqlDB, err := DataEngine.DB(); err != nil {
		return err
	} else {
		sqlDB.SetMaxIdleConns(DatabaseSetting.MaxIdleConns)
		sqlDB.SetMaxOpenConns(DatabaseSetting.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Hour)
	}
	return nil
}

// mysql database migrate
func NewDBMigrate() error {
	return nil
}
