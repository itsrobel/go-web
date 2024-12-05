package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/itsrobel/steflabs/internal/templates"
	"github.com/itsrobel/steflabs/internal/types"

	// "github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

var blogDir = "content/blog"

func HomeHandler(c *gin.Context) {
	blogDir, _ := os.ReadDir(blogDir)
	blogDirNames := make([]string, len(blogDir))
	for i, file := range blogDir {
		blogDirNames[i] = strings.Split(file.Name(), ".")[0]
	}
	fmt.Printf("Blog directory: %s\n", blogDirNames)

	bio, err := os.ReadFile("content/bio.md")

	var htmlBio string
	if err != nil {
		// TODO: I should make an actual page for this later c.String(http.StatusNotFound, "Page not found")
		htmlBio = "Page not found"
	} else {
		htmlBio = string(blackfriday.Run(bio,
			blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs),
		))
	}

	c.Writer.Header().Set("Blog-Type", "text/html")

	templates.Home(htmlBio, blogDirNames, getContacts()).Render(c.Request.Context(), c.Writer)
}

func AboutHandler(c *gin.Context) {
	about, err := os.ReadFile("content/about.md")

	var htmlBlog string
	if err != nil {
		// TODO: I should make an actual page for this later c.String(http.StatusNotFound, "Page not found")
		htmlBlog = "Page not found"
	} else {
		htmlBlog = string(blackfriday.Run(about,
			blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs),
		))
	}

	c.Writer.Header().Set("Blog-Type", "text/html")
	templates.About(htmlBlog).Render(c.Request.Context(), c.Writer)
}

func BlogHandler(c *gin.Context) {
	page := c.Param("page") + ".md"
	// blog, err := os.ReadFile(filepath.Join("blog", page))
	blog, err := os.ReadFile(filepath.Join(blogDir, page))
	var htmlBlog string
	if err != nil {
		// TODO: I should make an actual page for this later c.String(http.StatusNotFound, "Page not found")
		htmlBlog = "Page not found"
	} else {
		htmlBlog = string(blackfriday.Run(blog,
			blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs),
		))
	}

	// Render the template with the HTML blog
	c.Writer.Header().Set("Blog-Type", "text/html")
	// Pass the HTML blog as templ.HTML type
	templates.BlogPage(htmlBlog).Render(c.Request.Context(), c.Writer)
}

func getContacts() []types.Contact {
	contactInfo := []types.Contact{
		{Name: "Email", Icon: "fa-envelope", Link: "mailto:itsrobel.schwarz@gmail.com"},
		{Name: "GitHub", Icon: "fa-brands fa-github", Link: "https://github.com/itsrobel"},
		{Name: "LinkedIn", Icon: "fa-brands  fa-linkedin", Link: "https://www.linkedin.com/in/robel-schwarz/"},
	}
	return contactInfo
}

func RedirectSaveContact(c *gin.Context) {
	c.Header("HX-Redirect", "/save-contact")
	c.Status(200)
}

func SaveContact(c *gin.Context) {
	contact := struct {
		Fname string
		Lname string
		Phone string
		Email string
	}{
		Fname: "Robel",
		Lname: "Schwarz",
		Phone: "4257613775",
		Email: "itsrobel.schwarz@gmail.com",
	}

	vcard := fmt.Sprintf(
		"BEGIN:VCARD\nVERSION:3.0\nN:%s;%s;;;\nFN:%s %s\nTEL:%s\nEMAIL:%s\nEND:VCARD",
		contact.Lname, contact.Fname, contact.Fname, contact.Lname, contact.Phone, contact.Email,
	)

	c.Header("Blog-Type", "text/vcard; charset=utf-8")
	c.Header("Blog-Disposition", fmt.Sprintf("attachment; filename=\"%s.vcf\"", contact.Fname))
	c.String(http.StatusOK, vcard)
}
