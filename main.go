package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kartikkpawar/go-blog/database"
	"github.com/kartikkpawar/go-blog/router"
)

// init method gets called before main method
func init() {
	database.ConnectDB()
}

func main() {
	app := fiber.New()

	sqlDb, err := database.DBConnection.DB()
	if err != nil {
		panic("Error in SQL connection")
	}

	defer sqlDb.Close()

	app.Use(logger.New()) // To log the endpoints that are been hit

	router.SetupRoutes(app)

	app.Listen(":8000")
}
