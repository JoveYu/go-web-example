package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDatabase() error {
	var err error
	DB, err = gorm.Open(Config.DatabaseType, Config.Database)
	if err != nil {
		return err
	}

	return nil
}
