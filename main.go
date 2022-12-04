package main

import (
	"net/http"
	"time"
	"weight_management/database"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.DbOpen()
	db.AutoMigrate(&database.User{})
	db.AutoMigrate(&database.Weight{})
	engine := gin.Default()

	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "hello world",
		})
	})
	engine.GET("/users", func(c *gin.Context) {
		users := database.GetUsers()
		c.HTML(200, "users.html", gin.H{"users": users})
	})
	engine.POST("/users", func(c *gin.Context) {
		db.Create(&database.User{Name: "test", Age: 20, Birthday: time.Date(1999, 1, 1, 0, 0, 0, 0, time.Local)})
	})
	engine.Run(":3000")
}
