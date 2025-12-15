    package models

    import (
        "time"
        "github.com/google/uuid"
    )

type ClientHadist struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    ClientID  uuid.UUID `gorm:"type:uuid;not null;index" json:"client_id"`
    HadistID  uint      `gorm:"not null;index" json:"hadist_id"`
    Disabled  bool      `gorm:"default:true" json:"disabled"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

    func (ClientHadist) TableName() string {
        return "client_hadists"
    }