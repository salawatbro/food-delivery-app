package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
