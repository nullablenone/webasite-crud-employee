package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDataBase() *gorm.DB {
	dsn := "host=localhost user=postgres password=root dbname=crud-employee port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
