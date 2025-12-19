package services

import (
    "backend_masjid/internal/models"
    "backend_masjid/internal/repository"
    "fmt"
    "github.com/google/uuid"
    "backend_masjid/internal/cache"
)

type ClientHadistService struct {
    repo         *repository.ClientHadistRepository
    hadistRepo   *repository.HadistRepository
    cacheManager *cache.CacheManager
}

func NewClientHadistService(repo *repository.ClientHadistRepository, hadistRepo *repository.HadistRepository, cacheManager *cache.CacheManager) *ClientHadistService {
    return &ClientHadistService{repo: repo, hadistRepo: hadistRepo, cacheManager: cacheManager}
}

func (s *ClientHadistService) GetAllByClient(clientID string, limit, offset int) ([]models.HadistDTO, error) {
    cacheKey := fmt.Sprintf("client_hadist:%s:%d:%d", clientID, limit, offset)

    if data, ok := s.cacheManager.Get(cacheKey); ok {
        if hadists, ok := data.([]models.HadistDTO); ok {
            return hadists, nil
        }
    }

    hadists, err := s.hadistRepo.FindAll(limit, offset)
    if err != nil {
        return nil, err
    }

    disabled, err := s.repo.FindDisabledByClient(clientID)
    if err != nil {
        return nil, err
    }

    disabledMap := make(map[uint]bool)
    for _, d := range disabled {
        disabledMap[d.HadistID] = d.Disabled
    }

    var result []models.HadistDTO
    for _, h := range hadists {
        result = append(result, models.HadistDTO{
            ID:      h.ID,
            Konten:  h.Konten,
            Riwayat: h.Riwayat,
            Kitab:   h.Kitab,
            Enabled: !disabledMap[uint(h.ID)],
        })
    }

    s.cacheManager.Set(cacheKey, result)
    return result, nil
}

func (s *ClientHadistService) Disable(clientID string, hadistID uint) error {
    // ... validasi sama seperti sebelumnya
    err := s.repo.Disable(&models.ClientHadist{
        ClientID: uuid.MustParse(clientID),
        HadistID: hadistID,
        Disabled: true,
    })
    if err == nil {
        s.cacheManager.Invalidate("client_hadist:" + clientID)
    }
    return err
}

func (s *ClientHadistService) Enable(clientID string, hadistID uint) error {
    // ... validasi sama seperti sebelumnya
    err := s.repo.Enable(clientID, hadistID)
    if err == nil {
        s.cacheManager.Invalidate("client_hadist:" + clientID)
    }
    return err
}

func (s *ClientHadistService) SearchByKeyword(clientID, keyword string, limit, offset int) ([]models.HadistDTO, error) {
    // Query hadist global dengan keyword
    hadists, err := s.hadistRepo.Search(keyword, limit, offset)
    if err != nil {
        return nil, err
    }

    // Ambil data disabled untuk tenant
    disabled, err := s.repo.FindDisabledByClient(clientID)
    if err != nil {
        return nil, err
    }

    disabledMap := make(map[uint]bool)
    for _, d := range disabled {
        disabledMap[d.HadistID] = d.Disabled
    }

    // Gabungkan hasil
    var result []models.HadistDTO
    for _, h := range hadists {
        result = append(result, models.HadistDTO{
            ID:      h.ID,
            Konten:  h.Konten,
            Riwayat: h.Riwayat,
            Kitab:   h.Kitab,
            Enabled: !disabledMap[uint(h.ID)],
        })
    }

    return result, nil
}

func (s *ClientHadistService) Search(clientID, keyword string, limit, offset int) ([]models.HadistDTO, error) {
    cacheKey := fmt.Sprintf("client_hadist_search:%s:%s:%d:%d", clientID, keyword, limit, offset)

    // cek cache dulu
    if data, ok := s.cacheManager.Get(cacheKey); ok {
        if hadists, ok := data.([]models.HadistDTO); ok {
            return hadists, nil
        }
    }

    // query DB
    hadists, err := s.hadistRepo.Search(keyword, limit, offset)
    if err != nil {
        return nil, err
    }

    disabled, err := s.repo.FindDisabledByClient(clientID)
    if err != nil {
        return nil, err
    }

    disabledMap := make(map[uint]bool)
    for _, d := range disabled {
        disabledMap[d.HadistID] = d.Disabled
    }

    var result []models.HadistDTO
    for _, h := range hadists {
        result = append(result, models.HadistDTO{
            ID:      h.ID,
            Konten:  h.Konten,
            Riwayat: h.Riwayat,
            Kitab:   h.Kitab,
            Enabled: !disabledMap[uint(h.ID)],
        })
    }

    // simpan ke cache
    s.cacheManager.Set(cacheKey, result)

    return result, nil
}

func (s *ClientHadistService) GetStats(clientID string) (map[string]int64, error) {
    total, err := s.hadistRepo.CountAll()
    if err != nil {
        return nil, err
    }

    disabled, err := s.repo.CountDisabledByClient(clientID)
    if err != nil {
        return nil, err
    }

    enabled := total - disabled

    return map[string]int64{
        "total":    total,
        "enabled":  enabled,
        "disabled": disabled,
    }, nil
}