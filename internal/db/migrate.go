package db

import (
    "log"

    "gorm.io/gorm"
    "backend_masjid/internal/models"
)

// Migrate menjalankan auto-migrate untuk semua model GORM
func Migrate(db *gorm.DB) {
    err := db.AutoMigrate(
        &models.MasterClient{},
        &models.ClientBanner{},
        &models.ClientHadist{},
		&models.Hadist{},
		&models.AdminUser{},
		&models.DkmUser{},
        // tambahkan model lain di sini
    )
    if err != nil {
        log.Fatalf("❌ Gagal migrate schema: %v", err)
    }
    log.Println("✅ Schema berhasil di-migrate")
}