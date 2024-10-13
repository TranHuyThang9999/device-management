package pgsql

import (
	"device_management/common/log"
	"device_management/core/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPgsql() *gorm.DB {
	//dsn := "postgres://default:uIkv3Dg0fNAQ@ep-falling-sky-a41ji2wa-pooler.us-east-1.aws.neon.tech:5432/verceldb?sslmode=require"
	db, err := gorm.Open(postgres.Open(configs.Get().DataSource), &gorm.Config{})
	if err != nil {
		log.Fatal(err, "Failed to connect to the database")
		return nil
	}

	return db
}
