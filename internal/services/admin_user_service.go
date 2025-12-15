package services

import (
    "backend_masjid/internal/models"
    "backend_masjid/internal/repository"
)

type AdminUserService struct {
    repo *repository.AdminUserRepository
}

func NewAdminUserService(repo *repository.AdminUserRepository) *AdminUserService {
    return &AdminUserService{repo}
}

func (s *AdminUserService) Create(admin *models.AdminUser) error {
    return s.repo.Create(admin)
}

func (s *AdminUserService) GetByUsername(username string) (*models.AdminUser, error) {
    return s.repo.FindByUsername(username)
}

func (s *AdminUserService) Update(admin *models.AdminUser) error {
    return s.repo.Update(admin)
}

func (s *AdminUserService) Delete(id string) error {
    return s.repo.Delete(id)
}