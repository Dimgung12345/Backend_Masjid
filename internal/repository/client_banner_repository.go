package repository

import (
    "backend_masjid/internal/models"
    "gorm.io/gorm"
)

type BannerRepository struct {
    db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) *BannerRepository {
    return &BannerRepository{db}
}

// Create
func (r *BannerRepository) Create(banner *models.ClientBanner) error {
    return r.db.Create(banner).Error
}

func (r *BannerRepository) BulkInsert(banners []models.ClientBanner) error {
    return r.db.Create(&banners).Error
}

// Read All by ClientID
func (r *BannerRepository) FindAllByClient(clientID string) ([]models.ClientBanner, error) {
    var banners []models.ClientBanner
    err := r.db.Where("client_id = ?", clientID).Order("created_at desc").Find(&banners).Error
    return banners, err
}

// Read By ID
func (r *BannerRepository) FindByID(id int64) (*models.ClientBanner, error) {
    var banner models.ClientBanner
    err := r.db.First(&banner, id).Error
    if err != nil {
        return nil, err
    }
    return &banner, nil
}

// Update
func (r *BannerRepository) Update(banner *models.ClientBanner) error {
    return r.db.Save(banner).Error
}

// Delete
func (r *BannerRepository) Delete(id int64) error {
    return r.db.Delete(&models.ClientBanner{}, id).Error
}