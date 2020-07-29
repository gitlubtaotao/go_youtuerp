package global

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"youtuerp/conf"
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
	var err error
	DataEngine, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               conf.Configuration.DSN,
		DefaultStringSize: 256,
	}), &gorm.Config{
		Logger:      newLogger,
		PrepareStmt: true,
	})
	if err != nil {
		return err
	}
	sqlDB, err := DataEngine.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(1200)
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

// mysql database migrate
func NewDBMigrate() error {
	return nil
}
