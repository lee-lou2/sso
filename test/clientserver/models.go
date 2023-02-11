package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

var sqliteClient *gorm.DB

// AutoMigrate 마이그레이션
func AutoMigrate() {
	db := GetDatabase()
	db.AutoMigrate(
		&Config{},
	)
}

// GetDatabase 데이터베이스 연결
func GetDatabase() *gorm.DB {
	if sqliteClient == nil {
		// 데이터베이스 생성
		sqliteClient, _ = gorm.Open(sqlite.Open(os.Getenv("DATABASE_HOST_SQLITE")), &gorm.Config{})
	}
	return sqliteClient
}

// Config 기본 정보 저장
type Config struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Key       string `gorm:"unique;not null;type:varchar(50)" json:"key"`
	Value     string `gorm:"not null;type:varchar(255)" json:"value"`
}
