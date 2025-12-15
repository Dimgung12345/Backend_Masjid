package models

import (
    "time"
    "github.com/google/uuid"
)

type MasterClient struct {
    ID               uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
    Name             *string   `gorm:"size:100;not null" json:"name"`
    Location         *string   `gorm:"size:100" json:"location"`
    Timezone         *string   `gorm:"size:50;default:'Asia/Jakarta'" json:"timezone"`
    ConfigTitle      *string   `gorm:"size:100" json:"config_title"`
    ConfigBackground *string   `gorm:"size:100" json:"config_background"`
    ConfigSoundAlert *string   `gorm:"size:100" json:"config_sound_alert"`
    Logo             *string   `gorm:"size:100" json:"logo"`
    RunningText      *string   `gorm:"type:text" json:"running_text"`
    EnableHadis      *bool     `gorm:"default:true" json:"enable_hadis"`
    EnableHariBesar  *bool     `gorm:"default:true" json:"enable_hari_besar"`
    EnableKalender   *bool     `gorm:"default:true" json:"enable_kalender"`
    CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`

    LogoURL          string    `gorm:"-" json:"logo_url,omitempty"`
    BackgroundURL    string    `gorm:"-" json:"background_url,omitempty"`
    SoundURL         string    `gorm:"-" json:"sound_url,omitempty"`

}

func (MasterClient) TableName() string {
    return "master_client"
}