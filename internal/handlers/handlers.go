package handlers

import (
	"web/internal/templates"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	templates.Home().Render(c.Request.Context(), c.Writer)
}

func AboutHandler(c *gin.Context) {
	templates.About().Render(c.Request.Context(), c.Writer)
}
