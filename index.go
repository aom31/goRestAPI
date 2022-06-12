package main

import (
	"example/goRestAPI/model"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/go_orm_api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()
	r.GET("/users/read", func(c *gin.Context) {
		var users []model.User
		db.Find(&users)
		c.JSON(http.StatusOK, users)
	})
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user model.User
		db.First(&user, id)
		c.JSON(http.StatusOK, user)
	})

	r.POST("/user/create", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := db.Create(&user)
		c.JSON(200, gin.H{"RowsAffected": result.RowsAffected})

	})
	r.DELETE("/user/delete/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user model.User
		db.First(&user, id)
		db.Delete(&user)
	})
	r.PUT("/user/update", func(c *gin.Context) {
		var user model.User
		var updateUser model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.First(&updateUser, user.ID)
		updateUser.Username = user.Username
		updateUser.Fullname = user.Fullname
		updateUser.Password = user.Password
		updateUser.Email = user.Email

		db.Save(&updateUser)
		c.JSON(200, updateUser)
	})
	r.Use(cors.Default())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
