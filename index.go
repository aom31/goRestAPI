package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string ` json:"username"`
	Password string ` json:"password"`
}

func GetHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello golang"})
}
func Post(c *gin.Context) {
	var u User
	c.BindJSON(&u)
	c.JSON(http.StatusOK, gin.H{"user": u})
}
func PutPathParameter(c *gin.Context) {
	user := c.Param("user")
	c.JSON(http.StatusOK, gin.H{"message": user})
}

func DeleteQuery(c *gin.Context) {
	user := c.Query("user")
	c.JSON(http.StatusOK, gin.H{"message": user})
}
func main() {
	r := gin.Default()
	r.GET("/", GetHello)
	r.POST("/", Post)
	r.PUT("/path/:name", PutPathParameter)
	r.DELETE("/query", DeleteQuery)

	r.Run()
}
