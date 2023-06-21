package main

import (
	"go/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	handleRequests()
}

func handleRequests() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", routes.HomePageHandler)
	router.GET("/user/", routes.UserPageHandler)
	router.Run(":5000")
}
