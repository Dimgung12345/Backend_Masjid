package models

import (
    "time"
    "github.com/google/uuid"
)

type AdminUser struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Username  string    `gorm:"size:50;unique;not null"`
    Password  string    `gorm:"size:255;not null"`
    Role      string    `gorm:"size:20;not null;default:'admin'"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}
func (AdminUser) TableName() string {
    return "admin_user"
}
