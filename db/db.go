package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq"
    "stocky-assignment/config"
)

var DB *sql.DB

func Connect() {
    config.LoadEnv()
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        config.GetEnv("DB_HOST"), config.GetEnv("DB_PORT"),
        config.GetEnv("DB_USER"), config.GetEnv("DB_PASSWORD"),
        config.GetEnv("DB_NAME"))
    var err error
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Failed to open DB:", err)
    }
    err = DB.Ping()
    if err != nil {
        log.Fatal("Failed to connect DB:", err)
    }
    log.Println("✅ Connected to PostgreSQL")
}
