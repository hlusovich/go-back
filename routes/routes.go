package routes

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdmin bool   `json:"isAdmin"`
}

func HomePageHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "HOME",
	})
}

func UserPageHandler(c *gin.Context) {
	user := User{Name: "Miktia", Age: 27, IsAdmin: true}
	c.JSON(200, gin.H{
		"data": user,
	})
}
