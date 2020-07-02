package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config ...
type Config struct {
	DBHost     string
	DBDriver   string
	DBUser     string
	DBPassword string
	DBPort     string
	DBName     string
	APPPort    string
}

var defaultConfig = &Config{
	APPPort:    "8080",
	DBDriver:   "mysql",
	DBHost:     "127.0.0.1",
	DBName:     "fullstack_api",
	DBPassword: "admin",
	DBPort:     "3306",
	DBUser:     "root",
}

// GetConfig ...
func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		return defaultConfig
	}

	if os.Getenv("APP_PORT") != "" {
		defaultConfig.APPPort = os.Getenv("APP_PORT")
	}
	if os.Getenv("DB_DRIVER") != "" {
		defaultConfig.DBDriver = os.Getenv("DB_DRIVER")
	}
	if os.Getenv("DB_HOST") != "" {
		defaultConfig.DBHost = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_NAME") != "" {
		defaultConfig.DBName = os.Getenv("DB_NAME")
	}
	if os.Getenv("DB_PASSWORD") != "" {
		defaultConfig.DBPassword = os.Getenv("DB_PASSWORD")
	}
	if os.Getenv("DB_PORT") != "" {
		defaultConfig.DBPort = os.Getenv("DB_PORT")
	}
	if os.Getenv("DB_USER") != "" {
		defaultConfig.DBUser = os.Getenv("DB_USER")
	}
	return defaultConfig
}
