package models

import (
    "time"
    "github.com/google/uuid"
)

type ClientBanner struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    ClientID  uuid.UUID `gorm:"type:uuid;not null;index" json:"client_id"`
    Path      string    `json:"path"`
    CreatedAt time.Time `json:"created_at"`

    URL       string    `gorm:"-" json:"url,omitempty"`

}
func (ClientBanner) TableName() string {
    return "client_banners"
}
