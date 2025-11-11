package config

import (
    "os"

    "github.com/joho/godotenv"
)

func LoadEnv() {
    _ = godotenv.Load(".env")
}

func GetEnv(key string) string {
    return os.Getenv(key)
}
