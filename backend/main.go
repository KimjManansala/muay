package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/KimjManansala/muay/backend/pkg/db"
	"github.com/KimjManansala/muay/backend/pkg/handlers"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("../frontend/build", true)))
	router.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api") {
			fmt.Println("No route")
			c.File("../frontend/build/index.html")
		}
		//default 404 page not found
	})

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.POST("/addUser", h.AddUser)
		api.GET("/addUser", h.AddUser)
	}

	// Start and run the server
	router.Run(":4000")
}
