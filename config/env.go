package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost            string
	Port                  string
	DBUser                string
	DBPassword            string
	DBAddres              string
	DBName                string
	JWTExperetionInSecond int64
	JWTSecret             string
}

var Envs = initConfig()

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int64) int64 {
	if v, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:            getEnv("PUBLIC_HOST", "localhost"),
		Port:                  getEnv("PORT", "8080"),
		DBUser:                getEnv("DB_USER", "root"),
		DBPassword:            getEnv("DB_PASSWORD", ""),
		DBAddres:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName:                getEnv("DB_NAME", "trainDB"),
		JWTExperetionInSecond: getEnvInt("JWT_EXP", 3600*24*7),
		JWTSecret:             getEnv("JWT_SECRET", "not-secret-secret-anymore?"),
	}
}
