package repository

import (
    "backend_masjid/internal/models"
    "gorm.io/gorm"
)

type HariBesarRepository struct {
    db *gorm.DB
}

func NewHariBesarRepository(db *gorm.DB) *HariBesarRepository {
    return &HariBesarRepository{db}
}

// Create
func (r *HariBesarRepository) Create(hariBesar *models.HariBesar) error {
    return r.db.Create(hariBesar).Error
}

// Read All
func (r *HariBesarRepository) FindAll() ([]models.HariBesar, error) {
    var hariBesars []models.HariBesar
    err := r.db.Order("masehi asc").Find(&hariBesars).Error
    return hariBesars, err
}

// Read By Holiday Name
func (r *HariBesarRepository) FindByHoliday(name string) ([]models.HariBesar, error) {
    var hariBesars []models.HariBesar
    err := r.db.Where("holiday ILIKE ?", "%"+name+"%").Find(&hariBesars).Error
    return hariBesars, err
}

// Update
func (r *HariBesarRepository) Update(hariBesar *models.HariBesar) error {
    return r.db.Save(hariBesar).Error
}

// Delete
func (r *HariBesarRepository) Delete(id int64) error {
    return r.db.Delete(&models.HariBesar{}, id).Error
}