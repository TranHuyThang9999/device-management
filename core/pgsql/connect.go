package pgsql

import (
	"device_management/common/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPgsql() *gorm.DB {
	dsn := "postgres://default:uIkv3Dg0fNAQ@ep-falling-sky-a41ji2wa-pooler.us-east-1.aws.neon.tech:5432/verceldb?sslmode=require"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err, "Failed to connect to the database")
		return nil
	}

	return db
}
