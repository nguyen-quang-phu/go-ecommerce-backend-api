package initializes

import (
	"fmt"
	"time"

	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/global"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/internal/po"
	"github.com/nguyen-quang-phu/go-ecommerce-backend-api/pkg/settings"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDSN(config settings.MySQLSetting) string {
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	s := fmt.Sprintf(
		dsn,
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Dbname,
	)

	return s
}

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitDatabase() {
	config := global.Config.Mysql
	dsn := getDSN(config)
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{SkipDefaultTransaction: false},
	)
	checkErrorPanic(err, "initMysql initialization error")
	global.Logger.Info("initMysql initialization success")

	global.DB = db

	SetPoll(config)
	migrateTables()
}

func SetPoll(config settings.MySQLSetting) {
	sqlDB, err := global.DB.DB()
	if err != nil {
		fmt.Printf("mysql error: %s::", err)
	}

	sqlDB.SetConnMaxIdleTime(time.Duration(config.MaxIdleConns))
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime))
}

func migrateTables() {
	err := global.DB.AutoMigrate(&po.User{}, &po.Role{})
	if err != nil {
		fmt.Println("Migrating tables error:", err)
	}
}
