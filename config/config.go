package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort          string `json:"app_port"`
	PostgresUser     string `json:"postgres_user"`
	PostgresPassword string `json:"postgres_password"`
	PostgresHost     string `json:"postgres_host"`
	PostgresPort     string `json:"postgres_port"`
	PostgresDBName   string `json:"postgres_dbname"`
	MySQLUser        string `json:"mysql_user"`
	MySQLPassword    string `json:"mysql_password"`
	MySQLHost        string `json:"mysql_host"`
	MySQLDBName      string `json:"mysql_dbname"`
	JWTSecretKey     string `json:"jwt"`
}

func LoadConfig() (config *Config) {
	if err := godotenv.Load(os.ExpandEnv(".env")); err != nil {
		log.Printf("Error loading .env file: %s", err.Error())
		return nil
	}

	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8080"
	}

	return &Config{
		AppPort:          appPort,
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresDBName:   os.Getenv("POSTGRES_DBNAME"),
		MySQLUser:        os.Getenv("MYSQL_USER"),
		MySQLPassword:    os.Getenv("MYSQL_PASSWORD"),
		MySQLHost:        os.Getenv("MYSQL_HOST"),
		MySQLDBName:      os.Getenv("MYSQL_DBNAME"),
		JWTSecretKey:     os.Getenv("JWT_SECRET_KEY"),
	}
}
