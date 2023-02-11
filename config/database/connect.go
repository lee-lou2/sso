package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"sso/config/errors"
	"sso/config/errors/status"
	"strconv"
)

// 데이터베이스 기본 선언
var postgresClient *gorm.DB
var sqliteClient *gorm.DB

// GetDatabase 데이터베이스 연결
func GetDatabase() (*gorm.DB, error) {
	ProjectEnvironment := os.Getenv("PROJECT_ENVIRONMENT")
	if ProjectEnvironment == "LOCAL" {
		// Sqlite
		return connectSQLite()
	} else {
		// Postgres
		return connectPostgres()
	}
}

// connectSQLite SQLite 데이터베이스 연결
func connectSQLite() (*gorm.DB, error) {
	var err error
	if sqliteClient == nil {
		// 데이터베이스 생성
		sqliteClient, err = gorm.Open(
			sqlite.Open(os.Getenv("DATABASE_HOST_SQLITE")), &gorm.Config{},
		)
	}
	if err != nil {
		err = errors.New(status.FailedDBConnection, err)
	}
	return sqliteClient, err
}

// connectPostgres Postgres 데이터베이스 연결
func connectPostgres() (*gorm.DB, error) {
	var err error
	if postgresClient == nil {
		// 데이터베이스 생성
		p := os.Getenv("POSTGRES_PORT")
		port, err := strconv.ParseUint(p, 10, 32)
		if err != nil {
			panic("데이터베이스 포트 설정 실패")
		}
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DATABASE_HOST"),
			port,
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		)
		postgresClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		err = errors.New(status.FailedDBConnection, err)
	}
	return postgresClient, err
}
