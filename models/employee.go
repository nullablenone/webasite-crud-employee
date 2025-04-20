package models

import "gorm.io/gorm"

type Employee struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:40;not null"`
	Position string `gorm:"size:20;not null"`
	Email    string `gorm:"size:50;not null"`
	Phone    string `gorm:"size:15;not null"`
}

func MigrateEmployee(db *gorm.DB) error {
	err := db.AutoMigrate(&Employee{})
	return err
}
