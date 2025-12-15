package services

import (
	"backend_masjid/internal/models"
	"backend_masjid/internal/repository"
	"errors"

	"github.com/google/uuid"
)

type DkmUserService struct {
    repo *repository.DkmUserRepository
}

func NewDkmUserService(repo *repository.DkmUserRepository) *DkmUserService {
    return &DkmUserService{repo}
}

func (s *DkmUserService) Create(user *models.DkmUser) error {
    if user.ClientID == uuid.Nil || user.Username == "" || user.Password == "" {
        return errors.New("client_id, username, and password are required")
    }
    return s.repo.Create(user)
}

func (s *DkmUserService) GetAll() ([]models.DkmUser, error) {
    return s.repo.FindAll()
}

func (s *DkmUserService) GetByUsername(username string) (*models.DkmUser, error) {
    return s.repo.FindByUsername(username)
}

func (s *DkmUserService) Update(u *models.DkmUser) error {
    return s.repo.Update(u)
}

func (s *DkmUserService) Delete(id string) error {
    return s.repo.Delete(id)
}