package main

import (
	// "fmt"
	"os"
	"path/filepath"

	// web server support
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	// html + markdown processing

	"github.com/russross/blackfriday/v2"
)

func main() {
	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})

	// Load HTML templates
	r.LoadHTMLGlob("templates/*")

	// Serve static files
	r.Static("/static", "./static")

	// Handle root route
	r.GET("/", serveMarkdown("index.md"))

	// Handle other Markdown files
	r.GET("/:page", serveMarkdown(""))

	// Handle HTMX requests for page content
	r.GET("/content/:page", serveMarkdownContent)
	r.POST("/clicked", handleClicked)

	r.Run(":8080")
}

var clickCount = 0

// func handleClicked(c *gin.Context) {
// 	clickCount++
// 	c.String(http.StatusOK, fmt.Sprintf("Clicked %d times", clickCount))
// }

func serveMarkdown(defaultFile string) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.Param("page")
		if page == "" {
			page = defaultFile
		} else {
			page += ".md"
		}

		content, err := os.ReadFile(filepath.Join("content", page))
		if err != nil {
			c.String(http.StatusNotFound, "Page not found")
			return
		}

		// Convert Markdown to HTML
		html := blackfriday.Run(content)

		// Render the HTML within the layout template
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": string(html),
		})
	}
}

func serveMarkdownContent(c *gin.Context) {
	page := c.Param("page")
	content, err := os.ReadFile(filepath.Join("content", page+".md"))
	if err != nil {
		c.String(http.StatusNotFound, "Page not found")
		return
	}

	// Convert Markdown to HTML
	html := blackfriday.Run(content)

	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, string(html))
}
