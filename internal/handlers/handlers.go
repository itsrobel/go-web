package handlers

import (
	"log"
	"net/http"
	"os"
	"web/internal/templates"

	// "github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func HomeHandler(c *gin.Context) {
	templates.Home().Render(c.Request.Context(), c.Writer)
}

func AboutHandler(c *gin.Context) {
	templates.About().Render(c.Request.Context(), c.Writer)
}

func MarkdownHandler(c *gin.Context, file string) {
	// Path to your markdown file
	markdownFilePath := "string"
	content, err := os.ReadFile(markdownFilePath)
	if err != nil {
		log.Printf("Failed to read markdown file: %v", err)
		c.String(http.StatusInternalServerError, "Error loading markdown file")
		return
	}
	// Convert markdown to HTML
	// htmlContent := string(blackfriday.Run(content))

	htmlContent := string(blackfriday.Run(content,
		blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs),
	))

	// Render the template with the HTML content
	c.Writer.Header().Set("Content-Type", "text/html")

	// Pass the HTML content as templ.HTML type
	templates.MarkdownPage(htmlContent).Render(c.Request.Context(), c.Writer)
}
