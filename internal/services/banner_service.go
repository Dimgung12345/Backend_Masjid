package services

import (
    "backend_masjid/internal/models"
    "backend_masjid/internal/repository"
)

type BannerService struct {
    repo *repository.BannerRepository
}

func NewBannerService(repo *repository.BannerRepository) *BannerService {
    return &BannerService{repo}
}

func (s *BannerService) Create(b *models.ClientBanner) error {
    return s.repo.Create(b)
}

func (s *BannerService) GetAllByClient(clientID string) ([]models.ClientBanner, error) {
    return s.repo.FindAllByClient(clientID)
}

func (s *BannerService) GetByID(id int64) (*models.ClientBanner, error) {
    return s.repo.FindByID(id)
}

func (s *BannerService) Update(b *models.ClientBanner) error {
    return s.repo.Update(b)
}

func (s *BannerService) Delete(id int64) error {
    return s.repo.Delete(id)
}