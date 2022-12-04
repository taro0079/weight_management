package database

import "gorm.io/gorm"

type Weight struct {
	gorm.Model
	Weight float64
	UserID int
	User   User
}
