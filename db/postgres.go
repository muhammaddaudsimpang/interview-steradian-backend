package db

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Database string
	User     string
	Port     string
	Password string
	SSLMode  string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func InitDB() (*sql.DB, error) {
	config := DBConfig{
		Host: getEnv("HOST", "localhost"),
		Database: getEnv("DATABASE", "car_rental_db"),
		Port: getEnv("PORT", "5432"),
		User: getEnv("USER", "postgres"),
		Password: getEnv("PASSWORD", "123"),
		SSLMode: getEnv("SSLMOE", "disable"),
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Database,
		config.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err!= nil{
		return nil, err
	}

	fmt.Println("Database connect success")
	return db, nil
}
