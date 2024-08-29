package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/salawatbro/food-delivery-app/configs"
	"log"
)

func Connect(config configs.DatabaseConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Username, config.Password, config.DBName)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	log.Println("Connected to the PostgreSQL database successfully!")
	return db, nil
}

func Close(db *sqlx.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("Could not close the database connection: %v", err)
	}
}

func RunMigrate(db *sqlx.DB, action string) {
	var dir migrate.MigrationDirection
	migrations := &migrate.FileMigrationSource{
		Dir: "database/migrations",
	}
	if action == "up" {
		dir = migrate.Up
	} else if action == "down" {
		dir = migrate.Down
	} else {
		log.Fatalf("Invalid migration action: %s", action)
	}
	n, err := migrate.Exec(db.DB, "postgres", migrations, dir)
	if err != nil {
		log.Fatalf("Could not apply migrations: %v", err)
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
