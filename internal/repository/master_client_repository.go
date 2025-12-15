package repository

import (
    "backend_masjid/internal/models"
    "gorm.io/gorm"
)

type MasterClientRepository struct {
    db *gorm.DB
}

func NewMasterClientRepository(db *gorm.DB) *MasterClientRepository {
    return &MasterClientRepository{db}
}

func (r *MasterClientRepository) Create(client *models.MasterClient) error {
    return r.db.Create(client).Error
}

func (r *MasterClientRepository) FindAll() ([]models.MasterClient, error) {
    var clients []models.MasterClient
    err := r.db.Find(&clients).Error
    return clients, err
}

func (r *MasterClientRepository) FindByID(id string) (*models.MasterClient, error) {
    var client models.MasterClient
    err := r.db.First(&client, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &client, nil
}

func (r *MasterClientRepository) Update(client *models.MasterClient) error {
    return r.db.Model(&models.MasterClient{}).
        Where("id = ?", client.ID).
        Updates(map[string]interface{}{
            "name":              derefString(client.Name),
            "location":          derefString(client.Location),
            "timezone":          derefString(client.Timezone),
            "config_title":      derefString(client.ConfigTitle),
            "config_background": derefString(client.ConfigBackground),
            "config_sound_alert": derefString(client.ConfigSoundAlert),
            "logo":              derefString(client.Logo),
            "running_text":      derefString(client.RunningText),
            "enable_hadis":      derefBool(client.EnableHadis),
            "enable_hari_besar": derefBool(client.EnableHariBesar),
            "enable_kalender":   derefBool(client.EnableKalender),
        }).Error
}

func derefString(s *string) string {
    if s == nil {
        return ""
    }
    return *s
}

func derefBool(b *bool) bool {
    if b == nil {
        return false
    }
    return *b
}

func (r *MasterClientRepository) Delete(id string) error {
    return r.db.Delete(&models.MasterClient{}, "id = ?", id).Error
}