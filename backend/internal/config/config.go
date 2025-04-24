package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBuser             string
	DBpassword         string
	DBname             string
	DBmongoname        string
	DBcollection       string
	DBcollectionsecond string
}

func LoadConfig() *Config {
	// Загружаем переменные из .env файла
	err := godotenv.Load()
	if err != nil {
		log.Println("Файл .env не найден, используются переменные окружения")
	}

	return &Config{
		DBuser:             getEnv("DB_USER", ""),
		DBpassword:         getEnv("DB_PASSWORD", ""),
		DBname:             getEnv("DB_NAME", ""),
		DBmongoname:        getEnv("DB_MONGO_NAME", ""),
		DBcollection:       getEnv("DB_COLLECTION", ""),
		DBcollectionsecond: getEnv("DB_COLLECTION_SECOND", ""),
	}
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
