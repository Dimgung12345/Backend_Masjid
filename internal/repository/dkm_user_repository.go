package repository

import (
    "backend_masjid/internal/models"
    "gorm.io/gorm"
)

type DkmUserRepository struct {
    db *gorm.DB
}

func NewDkmUserRepository(db *gorm.DB) *DkmUserRepository {
    return &DkmUserRepository{db}
}

func (r *DkmUserRepository) Create(user *models.DkmUser) error {
    return r.db.Create(user).Error
}

func (r *DkmUserRepository) FindAll() ([]models.DkmUser, error) {
    var users []models.DkmUser
    err := r.db.Find(&users).Error
    return users, err
}

func (r *DkmUserRepository) FindByUsername(username string) (*models.DkmUser, error) {
    var user models.DkmUser
    err := r.db.Where("username = ?", username).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *DkmUserRepository) Update(user *models.DkmUser) error {
    return r.db.Save(user).Error
}

func (r *DkmUserRepository) Delete(id string) error {
    return r.db.Delete(&models.DkmUser{}, "id = ?", id).Error
}