package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "time"
)

// var clickCount = 0
//
//	func main() {
//		r := gin.Default()
//
//		// Serve static files
//		r.Static("/static", "./static")
//
//		r.LoadHTMLGlob("templates/*")
//		r.GET("/", handleIndex)
//		r.GET("/time", handleTime)
//		r.POST("/clicked", handleClicked)
//		fmt.Println("Server starting on :8080")
//		r.Run(":8080")
//	}
//
//	func handleIndex(c *gin.Context) {
//		c.HTML(http.StatusOK, "click.html", nil)
//	}
//
//	func handleTime(c *gin.Context) {
//		time.Sleep(1 * time.Second) // Simulate delay
//		c.String(http.StatusOK, time.Now().Format("15:04:05"))
//	}
func handleClicked(c *gin.Context) {
	clickCount++
	c.String(http.StatusOK, fmt.Sprintf("Clicked %d times", clickCount))
}
