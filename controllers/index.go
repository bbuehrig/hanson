package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Index shwos the Login-Site. If a User is logged in, than he will be
// redirected to the dashboard
//
func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"site":  "login",
		"title": "Login Site",
	})
}

type userData struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login handles login-data by POST-Request and logins userData
//
func Login(c *gin.Context) {

	// Bind userData
	var data userData
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// TODO: Login...
	// Delete Password
	data.Password = ""
	data.Username = "hanson"
	data.Email = "hanson@hanson.com"
	c.JSON(http.StatusOK, data)
}
