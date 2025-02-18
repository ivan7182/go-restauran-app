package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDB(dbAddress string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}

	return db
}
