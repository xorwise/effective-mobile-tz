package bootstrap

import (
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	APIPath          string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	PostgresHost     string
	PostgresPort     string
	LogLevel         slog.Level
	ListLimit        int
}

func ParseLogLeven(stringLevel string) (slog.Level, error) {
	var level slog.Level
	err := level.UnmarshalText([]byte(stringLevel))
	return level, err
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	logLevel, err := ParseLogLeven(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Fatal(err)
	}
	listLimit, err := strconv.Atoi(os.Getenv("LIST_LIMIT"))
	if err != nil {
		log.Fatal(err)
	}
	return &Env{
		APIPath:          os.Getenv("API_PATH"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDatabase: os.Getenv("POSTGRES_DATABASE"),
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		LogLevel:         logLevel,
		ListLimit:        listLimit,
	}
}
