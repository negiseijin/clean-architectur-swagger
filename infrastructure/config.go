package infrastructure

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB Env
}

type Env struct {
	Production  Dsn
	Development Dsn
}

type Dsn struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
	SslMode  string
	TimeZone string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}

	return &Config{
		DB: Env{
			Production: Dsn{
				Host:     os.Getenv("DATABASE_HOST"),
				User:     os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				DbName:   os.Getenv("POSTGRES_DB_NAME"),
				Port:     os.Getenv("POSTGRES_PORT"),
				SslMode:  os.Getenv("POSTGRES_SSLMODE"),
				TimeZone: os.Getenv("TZ"),
			},
			Development: Dsn{
				Host:     os.Getenv("DEV_DATABASE_HOST"),
				User:     os.Getenv("DEV_POSTGRES_USER"),
				Password: os.Getenv("DEV_POSTGRES_PASSWORD"),
				DbName:   os.Getenv("DEV_POSTGRES_DB_NAME"),
				Port:     os.Getenv("DEV_POSTGRES_PORT"),
				SslMode:  os.Getenv("DEV_POSTGRES_SSLMODE"),
				TimeZone: os.Getenv("DEV_TZ"),
			},
		},
	}
}
