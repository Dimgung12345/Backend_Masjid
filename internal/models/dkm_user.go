package models

import (
    "time"
    "github.com/google/uuid"
)

type DkmUser struct {
    ID        uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    ClientID  uuid.UUID    `gorm:"type:uuid;not null"`
    MasterClient MasterClient `gorm:"foreignKey:ClientID"`
    Username  string       `gorm:"size:50;unique;not null"`
    Password  string       `gorm:"size:255;not null"`
    Role      string       `gorm:"size:20;not null;default:'dkm'"`
    CreatedAt time.Time    `gorm:"autoCreateTime"`
}
func (DkmUser) TableName() string {
    return "dkm_user"
}