package models

type HadistDTO struct {
    ID      int64  `json:"id"`
    Konten  string `json:"konten"`
    Riwayat string `json:"riwayat"`
    Kitab   string `json:"kitab"`
    Enabled bool   `json:"enabled"`
}