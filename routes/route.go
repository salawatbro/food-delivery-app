package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Setup(app *fiber.App, db *sqlx.DB) {
	NewAuthRoutes(app, db)
	NewCourierRoutes(app, db)
	NewOrderRoutes(app, db)
	NewProductRoutes(app, db)
	NewAdminRoutes(app, db)
}
