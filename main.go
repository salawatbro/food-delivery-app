package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/salawatbro/food-delivery-app/configs"
	"github.com/salawatbro/food-delivery-app/database"
	"github.com/salawatbro/food-delivery-app/routes"
	"log"
)

func main() {
	config := configs.LoadConfig()
	db, err := database.Connect(config.DatabaseConfig)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			log.Fatalf("Could not close the database connection: %v", err)
		}
	}(db)

	app := fiber.New()

	routes.Setup(app, db)

	err = app.Listen(config.AppConfig.Host + ":" + config.AppConfig.Port)
	if err != nil {
		log.Fatalf("Could not listen to the server: %v", err)
	}
}
