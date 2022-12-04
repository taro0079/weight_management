package database

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func GetUsers() []User {
    db := DbOpen()
    var users []User
    db.Find(&users)
    return users
}
