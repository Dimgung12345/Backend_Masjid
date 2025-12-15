package models

import "time"

type JadwalSholat struct {
    ID        int64     `gorm:"primaryKey;autoIncrement"`
    Tanggal   time.Time `gorm:"unique;not null"`
    Lokasi    string    `gorm:"not null"`
    Subuh     string    `gorm:"not null"`
    Dzuhur    string    `gorm:"not null"`
    Ashar     string    `gorm:"not null"`
    Maghrib   string    `gorm:"not null"`
    Isya      string    `gorm:"not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}
func (JadwalSholat) TableName() string {
    return "jadwal_sholat_global"
}