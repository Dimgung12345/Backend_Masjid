package services

import (
    "backend_masjid/internal/models"
    "backend_masjid/internal/repository"
)

type HariBesarService struct {
    repo *repository.HariBesarRepository
}

func NewHariBesarService(repo *repository.HariBesarRepository) *HariBesarService {
    return &HariBesarService{repo}
}

func (s *HariBesarService) Create(h *models.HariBesar) error {
    return s.repo.Create(h)
}

func (s *HariBesarService) GetAll() ([]models.HariBesar, error) {
    return s.repo.FindAll()
}

func (s *HariBesarService) GetByHoliday(name string) ([]models.HariBesar, error) {
    return s.repo.FindByHoliday(name)
}

func (s *HariBesarService) Update(h *models.HariBesar) error {
    return s.repo.Update(h)
}

func (s *HariBesarService) Delete(id int64) error {
    return s.repo.Delete(id)
}