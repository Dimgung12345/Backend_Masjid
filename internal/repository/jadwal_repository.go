package repository

import (
    "backend_masjid/internal/models"
    "gorm.io/gorm"
)

type JadwalRepository struct {
    db *gorm.DB
}

func NewJadwalRepository(db *gorm.DB) *JadwalRepository {
    return &JadwalRepository{db}
}

func (r *JadwalRepository) Create(jadwal *models.JadwalSholat) error {
    return r.db.Create(jadwal).Error
}

func (r *JadwalRepository) FindAll() ([]models.JadwalSholat, error) {
    var jadwals []models.JadwalSholat
    err := r.db.Order("tanggal asc").Find(&jadwals).Error
    return jadwals, err
}

func (r *JadwalRepository) FindByDate(date string) (*models.JadwalSholat, error) {
    var jadwal models.JadwalSholat
    err := r.db.Where("tanggal = ?", date).First(&jadwal).Error
    if err != nil {
        return nil, err
    }
    return &jadwal, nil
}

func (r *JadwalRepository) FindByLokasi(lokasi string) (*models.JadwalSholat, error) {
    var jadlok models.JadwalSholat 
    err := r.db.Where("lokasi = ?", lokasi).First(&jadlok).Error
    if err != nil {
        return nil, err
    }
    return &jadlok, nil
}

func (r *JadwalRepository) Update(jadwal *models.JadwalSholat) error {
    return r.db.Save(jadwal).Error
}

func (r *JadwalRepository) Delete(id int64) error {
    return r.db.Delete(&models.JadwalSholat{}, id).Error
}