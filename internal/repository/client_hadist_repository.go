package repository

import (
    "backend_masjid/internal/models"
    "gorm.io/gorm"
)

type ClientHadistRepository struct {
    db *gorm.DB
}

func NewClientHadistRepository(db *gorm.DB) *ClientHadistRepository {
    return &ClientHadistRepository{db}
}

func (r *ClientHadistRepository) FindDisabledByClient(clientID string) ([]models.ClientHadist, error) {
    var disabled []models.ClientHadist
    err := r.db.Where("client_id = ? AND disabled = ?", clientID, true).Find(&disabled).Error
    return disabled, err
}

func (r *ClientHadistRepository) Disable(clientHadist *models.ClientHadist) error {
    return r.db.Save(clientHadist).Error
}

func (r *ClientHadistRepository) Enable(clientID string, hadistID uint) error {
    return r.db.Where("client_id = ? AND hadist_id = ?", clientID, hadistID).Delete(&models.ClientHadist{}).Error
}

func (r *HadistRepository) Search(keyword string, limit, offset int) ([]models.Hadist, error) {
    var hadists []models.Hadist
    like := "%" + keyword + "%"
    err := r.db.Where("konten LIKE ? OR riwayat LIKE ? OR kitab LIKE ?", like, like, like).
        Limit(limit).Offset(offset).Find(&hadists).Error
    return hadists, err
}