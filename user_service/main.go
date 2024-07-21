package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	database "superb.com/v0/user_service/db-config"
	routes "superb.com/v0/user_service/routes"
)

func main() {
	// Connect to the PostgreSQL database
	ConnectToDatabase()

	// Start the HTTP server
	StartHTTPServer()
}

func ConnectToDatabase() {
	database.ConnectToDB()
}

func StartHTTPServer() {
	app := fiber.New()
	routes.InitRoutes(app)
	port := os.Getenv("USER_SERVICE_PORT")
	app.Listen(":" + port)
}
