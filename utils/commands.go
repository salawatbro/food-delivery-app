package utils

import (
	"github.com/jmoiron/sqlx"
	"github.com/salawatbro/food-delivery-app/database"
	"os"
)

func SetupCommands(db *sqlx.DB) {
	for _, arg := range os.Args {
		if arg == "migrate --up" {
			database.RunMigrate(db, "up")
		} else if arg == "migrate --down" {
			database.RunMigrate(db, "down")
		}
	}
}
