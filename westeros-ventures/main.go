package main

import (
	"github.com/leif-runescribe/westeros-roster/routes"
	database "github.com/leif-runescribe/westeros-roster/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db := database.InitDB()

	// Set up the router
	r := gin.Default()

	// Set up routes
	routes.SetupUserRoutes(r, db)

	// Run the server
	r.Run(":8080")
}
