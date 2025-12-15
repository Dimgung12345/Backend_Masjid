package models

import "time"

type HariBesar struct {
    ID        int64     `gorm:"primaryKey;autoIncrement"`
    Masehi    time.Time `gorm:"not null"`
    Hijriyah  string    `gorm:"size:50;not null"`
    Holiday   string    `gorm:"size:120;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}
func (HariBesar) TableName() string {
    return "hari_besar_global"
}