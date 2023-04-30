package conf

import (
	"CyberAssetMapper/src/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDB() (*gorm.DB, error) {
	logMode := logger.Info
	if !viper.GetBool("mode.develop") {
		logMode = logger.Error
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: viper.GetString("db.dsn"), // DSN data source name
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true, //使用单数表名
		},
		Logger:                                   logger.Default.LogMode(logMode),
		DisableForeignKeyConstraintWhenMigrating: true, //禁用外键约束，使用逻辑外键
	})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConn")) //最大空闲连接数
	sqlDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConn")) //最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour)

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Task{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
