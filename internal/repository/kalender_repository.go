package repository

import (
    "backend_masjid/internal/models"
    "gorm.io/gorm"
)

type KalenderRepository struct {
    db *gorm.DB
}

func NewKalenderRepository(db *gorm.DB) *KalenderRepository {
    return &KalenderRepository{db}
}

// Create
func (r *KalenderRepository) Create(kalender *models.Kalender) error {
    return r.db.Create(kalender).Error
}

// Read All
func (r *KalenderRepository) FindAll() ([]models.Kalender, error) {
    var kalenders []models.Kalender
    err := r.db.Order("masehi asc").Find(&kalenders).Error
    return kalenders, err
}

// Read By Masehi Date
func (r *KalenderRepository) FindByDate(date string) (*models.Kalender, error) {
    var kalender models.Kalender
    err := r.db.Where("masehi = ?", date).First(&kalender).Error
    if err != nil {
        return nil, err
    }
    return &kalender, nil
}

// Update
func (r *KalenderRepository) Update(kalender *models.Kalender) error {
    return r.db.Save(kalender).Error
}

// Delete
func (r *KalenderRepository) Delete(id int64) error {
    return r.db.Delete(&models.Kalender{}, id).Error
}

// BulkCreate
func (r *KalenderRepository) BulkCreate(kalenders []models.Kalender, batchSize int) error {
    return r.db.CreateInBatches(kalenders, batchSize).Error
}