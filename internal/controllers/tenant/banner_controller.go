package tenant

import (
    "net/http"
    "strconv"

    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
)

type BannerController struct {
    service *services.BannerService
}

func NewBannerController(service *services.BannerService) *BannerController {
    return &BannerController{service}
}

// Tenant lihat semua banner miliknya
func (c *BannerController) GetAll(ctx *gin.Context) {
    clientID := ctx.GetString("client_id")
    if clientID == "" {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing client_id claim"})
        return
    }

    banners, err := c.service.GetAllByClient(clientID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    baseURL := "http://" + ctx.Request.Host
    for i := range banners {
        banners[i].URL = services.GetFileURL(baseURL, banners[i].Path)
    }

    ctx.JSON(http.StatusOK, banners)
}

// Tenant isi/ganti banner di slot tertentu
func (c *BannerController) Update(ctx *gin.Context) {
    clientID := ctx.GetString("client_id")
    if clientID == "" {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing client_id claim"})
        return
    }

    bannerID, _ := strconv.ParseInt(ctx.Param("bannerId"), 10, 64)

    file, err := ctx.FormFile("banner")
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "file banner wajib diupload"})
        return
    }

    path, err := services.SaveFile(ctx, file)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    banner, err := c.service.GetByID(bannerID) // pakai int64, sesuai service
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "banner not found"})
        return
    }

    // validasi kepemilikan banner
    if banner.ClientID.String() != clientID {
        ctx.JSON(http.StatusForbidden, gin.H{"error": "not your banner"})
        return
    }

    banner.Path = path
    if err := c.service.Update(banner); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "banner updated",
        "banner":  banner,
    })
}