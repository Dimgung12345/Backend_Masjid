package services

import (
    "errors"
    "backend_masjid/internal/models"
    "backend_masjid/internal/repository"
    "backend_masjid/internal/cache"
)

type MasterClientService struct {
    repo         *repository.MasterClientRepository
    cacheManager *cache.CacheManager
}

func NewMasterClientService(repo *repository.MasterClientRepository, cacheManager *cache.CacheManager) *MasterClientService {
    return &MasterClientService{repo: repo, cacheManager: cacheManager}
}

func (s *MasterClientService) Create(c *models.MasterClient) error {
    if c.Name == nil || *c.Name == "" ||
        c.Location == nil || *c.Location == "" ||
        c.Timezone == nil || *c.Timezone == "" {
        return errors.New("name, location, and timezone are required")
    }
    err := s.repo.Create(c)
    if err == nil {
        // invalidate cache list & client detail
        s.cacheManager.Invalidate("clients:all")
        s.cacheManager.Invalidate("client:" + c.ID.String())
    }
    return err
}

func (s *MasterClientService) GetAll() ([]models.MasterClient, error) {
    if data, ok := s.cacheManager.Get("clients:all"); ok {
        if clients, ok := data.([]models.MasterClient); ok {
            return clients, nil
        }
    }

    clients, err := s.repo.FindAll()
    if err != nil {
        return nil, err
    }

    s.cacheManager.Set("clients:all", clients)
    return clients, nil
}

func (s *MasterClientService) GetByID(id string) (*models.MasterClient, error) {
    if data, ok := s.cacheManager.Get("client:" + id); ok {
        if client, ok := data.(*models.MasterClient); ok {
            return client, nil
        }
    }

    client, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }

    s.cacheManager.Set("client:"+id, client)
    return client, nil
}

func (s *MasterClientService) Update(c *models.MasterClient) error {
    err := s.repo.Update(c)
    if err == nil {
        // invalidate cache
        s.cacheManager.Invalidate("client:" + c.ID.String())
        s.cacheManager.Invalidate("clients:all")
    }
    return err
}

func (s *MasterClientService) Delete(id string) error {
    err := s.repo.Delete(id)
    return err
}