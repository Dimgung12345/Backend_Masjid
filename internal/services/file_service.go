package services

import (
    "errors"
    "mime/multipart"
    "os"
    "path/filepath"
    "strings"
    "fmt"
    "github.com/gin-gonic/gin"
)

var (
    maxFileSize     = int64(5 * 1024 * 1024) // 5 MB
    allowedImageExt = []string{".jpg", ".jpeg", ".png"}
    allowedAudioExt = []string{".mp3", ".wav"}
)

// SaveFile menyimpan file ke storage dengan validasi ukuran & ekstensi
// return value: path relatif yang aman untuk disimpan di DB (contoh: "images/namafile.jpg")
func SaveFile(c *gin.Context, file *multipart.FileHeader) (string, error) {
    if file.Size > maxFileSize {
        return "", errors.New("file terlalu besar, maksimal 5MB")
    }

    ext := strings.ToLower(filepath.Ext(file.Filename))
    var folder string
    var relativeFolder string
    if contains(allowedImageExt, ext) {
        folder = "storage/images/"
        relativeFolder = "images/"
    } else if contains(allowedAudioExt, ext) {
        folder = "storage/audio/"
        relativeFolder = "audio/"
    } else {
        return "", errors.New("ekstensi file tidak diizinkan")
    }

    if err := os.MkdirAll(folder, os.ModePerm); err != nil {
        return "", err
    }

    fullPath := filepath.Join(folder, file.Filename)
    if err := c.SaveUploadedFile(file, fullPath); err != nil {
        return "", err
    }

    // 🔧 Simpan path relatif ke DB (sudah include folder, tanpa "storage/")
    cleanFileName := strings.ReplaceAll(file.Filename, "\\", "/")
    return relativeFolder + cleanFileName, nil
}

// Helper cek ekstensi
func contains(list []string, item string) bool {
    for _, v := range list {
        if v == item {
            return true
        }
    }
    return false
}

// GetFileURL bikin URL penuh dari path relatif
// contoh: baseURL="http://192.168.1.7:8080", path="images/namafile.jpg"
// hasil: "http://192.168.1.7:8080/storage/images/namafile.jpg"
func GetFileURL(baseURL, path string) string {
    cleanPath := strings.ReplaceAll(path, "\\", "/") // normalisasi backslash
    cleanPath = strings.TrimPrefix(cleanPath, "/")   // buang leading slash
    return fmt.Sprintf("%s/storage/%s", baseURL, cleanPath)
}