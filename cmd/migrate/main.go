package main

import (
    "log"

    "backend_masjid/internal/config"
    "backend_masjid/internal/db"
)

func main() {
    // 1. Load config dari env/file
    config.LoadConfig()

    // 2. Init koneksi DB
    db.InitDB()

    // 3. Jalankan migrasi schema
    db.Migrate(db.DB)

    log.Println("✅ Migrasi selesai, schema sudah sync dengan model GORM")
}