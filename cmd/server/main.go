package main

import (
	"web/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/", handlers.HomeHandler)
	r.GET("/about", handlers.AboutHandler)

	r.Run(":8080")
}
