package services

import (
    "backend_masjid/internal/models"
    "backend_masjid/internal/repository"
)

type JadwalService struct {
    repo *repository.JadwalRepository
}

func NewJadwalService(repo *repository.JadwalRepository) *JadwalService {
    return &JadwalService{repo}
}

// Create
func (s *JadwalService) Create(j *models.JadwalSholat) error {
    return s.repo.Create(j)
}

// Read All
func (s *JadwalService) GetAll() ([]models.JadwalSholat, error) {
    return s.repo.FindAll()
}

// Read By Date
func (s *JadwalService) GetByDate(date string) (*models.JadwalSholat, error) {
    return s.repo.FindByDate(date)
}

func (s *JadwalService) GetByLokasi(lokasi string) (*models.JadwalSholat, error) {
    return s.repo.FindByLokasi(lokasi)
}

// Update
func (s *JadwalService) Update(j *models.JadwalSholat) error {
    return s.repo.Update(j)
}

// Delete
func (s *JadwalService) Delete(id int64) error {
    return s.repo.Delete(id)
}