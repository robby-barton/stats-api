package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBParams struct {
	Host     string
	Port     int64
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewDatabase(params *DBParams) (*gorm.DB, error) {
	if params != nil {
		return postgresDB(params)
	}
	return sqliteDB()
}

func postgresDB(params *DBParams) (*gorm.DB, error) {
	connInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		params.Host,
		params.Port,
		params.User,
		params.Password,
		params.DBName,
		params.SSLMode,
	)

	return gorm.Open(postgres.Open(connInfo), &gorm.Config{
		SkipDefaultTransaction: true, // handle my own transactions
	})
}

func sqliteDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("db/cfb.db"), &gorm.Config{
		SkipDefaultTransaction: true, // handle my own transactions
	})
}
