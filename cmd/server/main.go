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
	r.GET("/:page", handlers.MarkdownHandler)

	// TODO: for the files found in the content folder server them in content
	// r.GET("/markdown", handlers.MarkdownHandler)

	r.Run(":8080")
}
