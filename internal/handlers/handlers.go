package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"web/internal/templates"
	"web/internal/types"

	// "github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func HomeHandler(c *gin.Context) {
	contentDir, _ := os.ReadDir("content")

	contentDirNames := make([]string, len(contentDir))
	for i, file := range contentDir {
		contentDirNames[i] = strings.Split(file.Name(), ".")[0]
	}

	fmt.Printf("Content directory: %s\n", contentDirNames)

	templates.Home(contentDirNames, getContacts()).Render(c.Request.Context(), c.Writer)
}

func AboutHandler(c *gin.Context) {
	templates.About().Render(c.Request.Context(), c.Writer)
}

func ContentHandler(c *gin.Context) {
	page := c.Param("page") + ".md"
	// content, err := os.ReadFile(filepath.Join("content", page))
	content, err := os.ReadFile(filepath.Join("content", page))
	var htmlContent string
	if err != nil {
		// TODO: I should make an actual page for this later c.String(http.StatusNotFound, "Page not found")
		htmlContent = "Page not found"
	} else {
		htmlContent = string(blackfriday.Run(content,
			blackfriday.WithExtensions(blackfriday.CommonExtensions|blackfriday.AutoHeadingIDs),
		))
	}

	// Render the template with the HTML content
	c.Writer.Header().Set("Content-Type", "text/html")
	// Pass the HTML content as templ.HTML type
	templates.ContentPage(htmlContent).Render(c.Request.Context(), c.Writer)
}

func getContacts() []types.Contact {
	contactInfo := []types.Contact{
		{Name: "Email", Icon: "fa-envelope", Link: "mailto:itsrobel.schwarz@gmail.com"},
		{Name: "GitHub", Icon: "fa-github", Link: "https://github.com/itsrobel"},
		{Name: "LinkedIn", Icon: "fa-linkedin", Link: "https://www.linkedin.com/in/robel-schwarz/"},
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

	c.Header("Content-Type", "text/vcard; charset=utf-8")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.vcf\"", contact.Fname))
	c.String(http.StatusOK, vcard)
}
