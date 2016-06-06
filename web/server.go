// Package web provides the Gin-Framework initialization. It declares all
// routes and the endpoints.
package web

import (
	"hanson/controllers"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

// StartServer initialzes the routes, defines the endpoints and start the
// Web-Server.
//
// PARAMETERS:
//    portNumber - Portnumber the webserver is listening
//		debugMode  - if true debug mode ist on - Webserver prints out more informations
//
func StartServer(portNumber string, debugMode bool) {
	log.Info("Initialize Web-Server...")

	// Set Debug/ReleaseMode
	if debugMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup Gin
	router := gin.Default()

	// Route definitions
	router.Static("/js", "./public/js")   // JS-Route
	router.Static("/css", "./public/css") // CSS-Route

	// Set Template Folder
	router.LoadHTMLGlob("views/*")

	// Websocket-Handler
	router.GET("/ping", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	// ########## Routes ##########

	// ____ GET ____
	// Index / Login
	router.GET("/", controllers.Index) // login

	// Ping
	router.GET("/dashboard", controllers.Dashboard) // Dashboard

	// ____ POST ____
	// Login
	router.POST("/login", controllers.Login)

	// Start Webserver
	log.Info("Start Web-Server...")
	router.Run(":" + portNumber)
}

// Initialize Gorilla Websocket-Upgrader
//
var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// wshandler is the Websocket-handler called by gin
//
func wshandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Websocket-connection from %s", r.UserAgent())

	// Create Connection
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("Failed to set websocket upgrade: %v", err)
		return
	}

	// Sample....
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		log.Infof("Message received: %s", msg)
		if string(msg) == "ping" {
			sendmsg := []byte("pong")

			conn.WriteMessage(t, sendmsg)
		} else {
			conn.WriteMessage(t, msg)
		}
	}
}
