package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//
func Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"site":  "dashboard",
		"title": "Dashboard",
	})
}
