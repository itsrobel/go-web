package main

import (
	// "fmt"
	"os"
	"path/filepath"

	// web server support
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func main() {
	watcher()
	server()
}

func server() {
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

	// Handle HTMX requests for page blog
	r.POST("/clicked", handleClicked)
	println("Server running on port 8080")

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

		blog, err := os.ReadFile(filepath.Join("blog", page))
		if err != nil {
			c.String(http.StatusNotFound, "Page not found")
			return
		}

		// Convert Markdown to HTML
		html := blackfriday.Run(blog)

		// Render the HTML within the layout template
		c.HTML(http.StatusOK, "index.html", gin.H{
			"blog": string(html),
		})
	}
}

func serveMarkdownBlog(c *gin.Context) {
	page := c.Param("page")
	blog, err := os.ReadFile(filepath.Join("blog", page+".md"))
	if err != nil {
		c.String(http.StatusNotFound, "Page not found")
		return
	}

	// Convert Markdown to HTML
	html := blackfriday.Run(blog)

	c.Header("Blog-Type", "text/html")
	c.String(http.StatusOK, string(html))
}
