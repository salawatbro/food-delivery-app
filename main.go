package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salawatbro/food-delivery-app/configs"
	"github.com/salawatbro/food-delivery-app/database"
	"github.com/salawatbro/food-delivery-app/routes"
	"github.com/salawatbro/food-delivery-app/utils"
	"log"
)

func main() {
	config := configs.LoadConfig()
	db, err := database.Connect(config.DatabaseConfig)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	utils.SetupCommands(db)
	defer database.Close(db)

	app := fiber.New()

	routes.Setup(app, db)

	err = app.Listen(config.AppConfig.Host + ":" + config.AppConfig.Port)
	if err != nil {
		log.Fatalf("Could not listen to the server: %v", err)
	}
}
