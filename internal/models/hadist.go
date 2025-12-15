package models

import "time"

type Hadist struct {
    ID        int64     `gorm:"primaryKey;autoIncrement"`
    Konten    string    `gorm:"type:text;not null"`
    Riwayat   string    `gorm:"size:120"`
    Kitab     string    `gorm:"size:120"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}
func (Hadist) TableName() string {
    return "hadist_global"
}