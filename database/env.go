package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func GetDbConfig() *DbConfig {
	godotenv.Load()

	portAsString := os.Getenv("DATABASE_PORT")
	port, err := strconv.Atoi(portAsString)
	if err != nil {
		log.Fatal("could not get datbase port")
	}
	dbConfig := DbConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     port,
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		DbName:   os.Getenv("DATABASE_DBNAME"),
	}

	return &dbConfig
}

func (dbConfig DbConfig) getConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DbName)

}
