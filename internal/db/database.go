package db

import (
    "fmt"
    "log"

    "backend_masjid/internal/config"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        config.Cfg.DbHost,
        config.Cfg.DbUser,
        config.Cfg.DbPass,
        config.Cfg.DbName,
        config.Cfg.DbPort,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    DB = db
    log.Println("Database connected successfully")
}