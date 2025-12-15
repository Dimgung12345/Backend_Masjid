package repository

import (
    "backend_masjid/internal/models"
    "gorm.io/gorm"
)

type AdminUserRepository struct {
    db *gorm.DB
}

func NewAdminUserRepository(db *gorm.DB) *AdminUserRepository {
    return &AdminUserRepository{db}
}

func (r *AdminUserRepository) Create(admin *models.AdminUser) error {
    return r.db.Create(admin).Error
}

func (r *AdminUserRepository) FindByUsername(username string) (*models.AdminUser, error) {
    var admin models.AdminUser
    err := r.db.Where("username = ?", username).First(&admin).Error
    if err != nil {
        return nil, err
    }
    return &admin, nil
}

func (r *AdminUserRepository) Update(admin *models.AdminUser) error {
    return r.db.Save(admin).Error
}

func (r *AdminUserRepository) Delete(id string) error {
    return r.db.Delete(&models.AdminUser{}, "id = ?", id).Error
}