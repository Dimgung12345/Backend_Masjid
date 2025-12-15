package services

import (
    "backend_masjid/internal/models"
    "backend_masjid/internal/repository"
)

type KalenderService struct {
    repo *repository.KalenderRepository
}

func NewKalenderService(repo *repository.KalenderRepository) *KalenderService {
    return &KalenderService{repo}
}

func (s *KalenderService) Create(k *models.Kalender) error {
    return s.repo.Create(k)
}

func (s *KalenderService) GetAll() ([]models.Kalender, error) {
    return s.repo.FindAll()
}

func (s *KalenderService) GetByDate(date string) (*models.Kalender, error) {
    return s.repo.FindByDate(date)
}

func (s *KalenderService) Update(k *models.Kalender) error {
    return s.repo.Update(k)
}

func (s *KalenderService) Delete(id int64) error {
    return s.repo.Delete(id)
}

func (s *KalenderService) BulkCreate(kalenders []models.Kalender, batchSize int) error {
    return s.repo.BulkCreate(kalenders, batchSize)
}