package services

import (
    "backend_masjid/internal/models"
    "backend_masjid/internal/repository"
    "backend_masjid/internal/cache"
    "fmt"
    "errors"
)

type BannerService struct {
    repo         *repository.BannerRepository
    cacheManager *cache.CacheManager
}

func NewBannerService(repo *repository.BannerRepository, cacheManager *cache.CacheManager) *BannerService {
    return &BannerService{repo: repo, cacheManager: cacheManager}
}

func (s *BannerService) Create(b *models.ClientBanner) error {
    err := s.repo.Create(b)
    if err == nil {
        s.cacheManager.Invalidate("banner:" + b.ClientID.String())
    }
    return err
}

func (s *BannerService) BulkCreate(banners []models.ClientBanner) error {
    if len(banners) == 0 {
        return errors.New("no banners provided")
    }

    // Insert ke DB
    err := s.repo.BulkInsert(banners)
    if err == nil {
        // invalidate cache client
        s.cacheManager.Invalidate("banner:" + banners[0].ClientID.String())
    }
    return err
}

func (s *BannerService) GetAllByClient(clientID string) ([]models.ClientBanner, error) {
    cacheKey := "banner:" + clientID
    if data, ok := s.cacheManager.Get(cacheKey); ok {
        if banners, ok := data.([]models.ClientBanner); ok {
            return banners, nil
        }
    }

    banners, err := s.repo.FindAllByClient(clientID)
    if err != nil {
        return nil, err
    }

    s.cacheManager.Set(cacheKey, banners)
    return banners, nil
}

func (s *BannerService) GetByID(id int64) (*models.ClientBanner, error) {
    cacheKey := fmt.Sprintf("banner:id:%d", id)
    if data, ok := s.cacheManager.Get(cacheKey); ok {
        if banner, ok := data.(*models.ClientBanner); ok {
            return banner, nil
        }
    }

    banner, err := s.repo.FindByID(id)
    if err != nil {
        return nil, err
    }

    s.cacheManager.Set(cacheKey, banner)
    return banner, nil
}

func (s *BannerService) Update(b *models.ClientBanner) error {
    err := s.repo.Update(b)
    if err == nil {
        s.cacheManager.Invalidate("banner:" + b.ClientID.String())
        s.cacheManager.Invalidate(fmt.Sprintf("banner:id:%d", b.ID))
    }
    return err
}

func (s *BannerService) Delete(id int64, clientID string) error {
    err := s.repo.Delete(id)
    if err == nil {
        s.cacheManager.Invalidate("banner:" + clientID)
        s.cacheManager.Invalidate(fmt.Sprintf("banner:id:%d", id))
    }
    return err
}