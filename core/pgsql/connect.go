package pgsql

import (
	"device_management/common/log"
	"device_management/core/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPgsql() *gorm.DB {
	db, err := gorm.Open(postgres.Open(configs.Get().DataSource), &gorm.Config{})
	if err != nil {
		log.Fatal(err, "error connect pgsql")
		return nil
	}
	return db
}
