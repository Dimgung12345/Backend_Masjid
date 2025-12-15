package services

import (
    "backend_masjid/internal/models"
    "backend_masjid/internal/repository"
    "fmt"
    "github.com/google/uuid"
)

type ClientHadistService struct {
    repo       *repository.ClientHadistRepository
    hadistRepo *repository.HadistRepository
}

func NewClientHadistService(repo *repository.ClientHadistRepository, hadistRepo *repository.HadistRepository) *ClientHadistService {
    return &ClientHadistService{repo: repo, hadistRepo: hadistRepo}
}

// Ambil semua hadist global + status enabled/disabled untuk tenant
func (s *ClientHadistService) GetAllByClient(clientID string, limit, offset int) ([]models.HadistDTO, error) {
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

    return result, nil
}

// Disable hadist untuk tenant
func (s *ClientHadistService) Disable(clientID string, hadistID uint) error {
    if hadistID == 0 {
        return fmt.Errorf("invalid hadist_id")
    }

    exists, err := s.hadistRepo.Exists(hadistID)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("hadist %d not found", hadistID)
    }

    ch := models.ClientHadist{
        ClientID: uuid.MustParse(clientID),
        HadistID: hadistID,
        Disabled: true,
    }
    return s.repo.Disable(&ch)
}

// Enable hadist untuk tenant
func (s *ClientHadistService) Enable(clientID string, hadistID uint) error {
    if hadistID == 0 {
        return fmt.Errorf("invalid hadist_id")
    }

    exists, err := s.hadistRepo.Exists(hadistID)
    if err != nil {
        return err
    }
    if !exists {
        return fmt.Errorf("hadist %d not found", hadistID)
    }

    return s.repo.Enable(clientID, hadistID)
}