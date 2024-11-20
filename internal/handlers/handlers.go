package handlers

import (
	"net/http"
	"os"
	"path/filepath"
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

func MarkdownHandler(c *gin.Context) {
	page := c.Param("page") + ".md"
	// content, err := os.ReadFile(filepath.Join("content", page))
	content, err := os.ReadFile(filepath.Join("content", page))
	if err != nil {
		c.String(http.StatusNotFound, "Page not found")
	}
	htmlContent := string(blackfriday.Run(content,
		blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs),
	))

	// Render the template with the HTML content
	c.Writer.Header().Set("Content-Type", "text/html")
	// Pass the HTML content as templ.HTML type
	templates.MarkdownPage(htmlContent).Render(c.Request.Context(), c.Writer)
}
