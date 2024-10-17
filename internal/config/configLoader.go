package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	DbHost       string `map-structure:"DB_HOST"`
	DbPort       string `map-structure:"DB_PORT"`
	DbName       string `map-structure:"DB_NAME"`
	DbUser       string `map-structure:"DB_USER"`
	DbPass       string `map-structure:"DB_PASS"`
	DbSslMode    string `map-structure:"DB_SSLMODE"`
	AllowOrigins string `map-structure:"ALLOW_ORIGINS"`
	AppPort      string `map-structure:"APP_PORT"`
	Url          string `map-structure:"URL"`
	ApiKey       string `map-structure:"API_KEY"`
}

var conf Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Warning("Couldn't load env variables to connect to DB...")
	}

	conf = Config{
		AppPort:      getEnv("APP_PORT", ""),
		Url:          getEnv("URL", ""),
		ApiKey:       getEnv("API_KEY", ""),
		DbHost:       getEnv("DB_HOST", ""),
		DbPort:       getEnv("DB_PORT", ""),
		DbName:       getEnv("DB_NAME", ""),
		DbUser:       getEnv("DB_USER", ""),
		DbPass:       getEnv("DB_PASS", ""),
		DbSslMode:    getEnv("DB_SSLMODE", ""),
		AllowOrigins: getEnv("ALLOW_ORIGINS", ""),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func GetConfigurationInstance() Config {
	return conf
}
