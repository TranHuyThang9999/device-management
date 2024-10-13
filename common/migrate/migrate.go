package migrate

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MigratePgsqlS() {
	host := "ep-falling-sky-a41ji2wa-pooler.us-east-1.aws.neon.tech"
	port := "5432"
	user := "default"
	password := "uIkv3Dg0fNAQ"
	dbname := "verceldb"
	sslmode := "require"

	// Tạo chuỗi kết nối
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslmode)

	// Mở kết nối đến cơ sở dữ liệu
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Đọc tệp SQL
	sqlFile, err := os.ReadFile("schema/schema.sql")
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}

	// Chạy các câu lệnh SQL từ tệp
	if err := db.Exec(string(sqlFile)).Error; err != nil {
		log.Fatalf("Failed to execute SQL commands: %v", err)
	}

	log.Println("Database migrated successfully!")
}
