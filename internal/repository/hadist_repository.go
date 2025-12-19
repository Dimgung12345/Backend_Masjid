package repository

import (
    "backend_masjid/internal/models"
    "gorm.io/gorm"
)

type HadistRepository struct {
    db *gorm.DB
}

func NewHadistRepository(db *gorm.DB) *HadistRepository {
    return &HadistRepository{db}
}

// Create
func (r *HadistRepository) Create(hadist *models.Hadist) error {
    return r.db.Create(hadist).Error
}

// Read All
func (r *HadistRepository) FindAll(limit, offset int) ([]models.Hadist, error) {
    var hadists []models.Hadist
    err := r.db.
        Order("created_at desc").
        Limit(limit).
        Offset(offset).
        Find(&hadists).Error
    return hadists, err
}


// Read By ID
func (r *HadistRepository) FindByID(id int64) (*models.Hadist, error) {
    var hadist models.Hadist
    err := r.db.First(&hadist, id).Error
    if err != nil {
        return nil, err
    }
    return &hadist, nil
}

// Update
func (r *HadistRepository) Update(hadist *models.Hadist) error {
    return r.db.Save(hadist).Error
}

// Delete
func (r *HadistRepository) Delete(id int64) error {
    return r.db.Delete(&models.Hadist{}, id).Error
}

func (r *HadistRepository) Exists(hadistID uint) (bool, error) {
    var count int64
    err := r.db.Model(&models.Hadist{}).Where("id = ?", hadistID).Count(&count).Error
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

func (r *HadistRepository) CountAll() (int64, error) {
    var count int64
    err := r.db.Model(&models.Hadist{}).Count(&count).Error
    return count, err
}
