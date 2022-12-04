package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbOpen() *gorm.DB {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("error")
	}
	return db
}
