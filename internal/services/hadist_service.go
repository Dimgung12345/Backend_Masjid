package services

import (
    "backend_masjid/internal/models"
    "backend_masjid/internal/repository"
)

type HadistService struct {
    repo *repository.HadistRepository
}

func NewHadistService(repo *repository.HadistRepository) *HadistService {
    return &HadistService{repo}
}

func (s *HadistService) Create(h *models.Hadist) error {
    return s.repo.Create(h)
}

func (s *HadistService) GetAll(limit, offset int) ([]models.Hadist, error) {
    return s.repo.FindAll(limit, offset)
}


func (s *HadistService) GetByID(id int64) (*models.Hadist, error) {
    return s.repo.FindByID(id)
}

func (s *HadistService) Update(h *models.Hadist) error {
    return s.repo.Update(h)
}

func (s *HadistService) Delete(id int64) error {
    return s.repo.Delete(id)
}