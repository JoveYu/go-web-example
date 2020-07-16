package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"test.com/example/common/middleware"
)

var DB *gorm.DB

func InitDatabase() error {
	// 加载自定义logger
	logger := middleware.NewGORMLogger()

	var err error
	DB, err = gorm.Open(Config.DatabaseType, Config.Database)
	DB.SetLogger(logger)
	DB.LogMode(true)
	if err != nil {
		return err
	}

	return nil
}
