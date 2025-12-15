package models

import "time"

type Kalender struct {
    ID        int64     `gorm:"primaryKey;autoIncrement"`
    Masehi    time.Time `gorm:"unique;not null"`
    Hijriyah  string    `gorm:"size:50;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}
func (Kalender) TableName() string {
    return "kalender"
}