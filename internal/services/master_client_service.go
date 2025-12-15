package services

import (
    "errors"
    "backend_masjid/internal/models"
    "backend_masjid/internal/repository"
)

type MasterClientService struct {
    repo *repository.MasterClientRepository
}

func NewMasterClientService(repo *repository.MasterClientRepository) *MasterClientService {
    return &MasterClientService{repo}
}

func (s *MasterClientService) Create(c *models.MasterClient) error {
    if c.Name == nil || *c.Name == "" ||
       c.Location == nil || *c.Location == "" ||
       c.Timezone == nil || *c.Timezone == "" {
        return errors.New("name, location, and timezone are required")
    }
    return s.repo.Create(c)
}

func (s *MasterClientService) GetAll() ([]models.MasterClient, error) {
    return s.repo.FindAll()
}

func (s *MasterClientService) GetByID(id string) (*models.MasterClient, error) {
    return s.repo.FindByID(id)
}

func (s *MasterClientService) Update(c *models.MasterClient) error {
    return s.repo.Update(c)
}

func (s *MasterClientService) Delete(id string) error {
    return s.repo.Delete(id)
}