package admin

import (
    "net/http"
    "strconv"
    "backend_masjid/internal/models"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

type BannerController struct {
    service *services.BannerService
}

func NewBannerController(service *services.BannerService) *BannerController {
    return &BannerController{service}
}

// POST /admin/banners
func (c *BannerController) Create(ctx *gin.Context) {
    clientID := ctx.PostForm("client_id")
    if clientID == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "client_id required"})
        return
    }
    clientUUID := uuid.MustParse(clientID)

    form, err := ctx.MultipartForm()
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid multipart form"})
        return
    }

    files := form.File["banner"]
    if len(files) == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "banner file required"})
        return
    }

    // ✅ Single create
    if len(files) == 1 {
        path, err := services.SaveFile(ctx, files[0])
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        banner := models.ClientBanner{
            ClientID: clientUUID,
            Path:     path,
        }

        if err := c.service.Create(&banner); err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // selalu return array agar FE konsisten
        ctx.JSON(http.StatusCreated, []models.ClientBanner{banner})
        return
    }

    // ✅ Bulk create
    if len(files) > 5 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "maksimal 5 banner"})
        return
    }

    var banners []models.ClientBanner
    for _, file := range files {
        path, err := services.SaveFile(ctx, file)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        banners = append(banners, models.ClientBanner{
            ClientID: clientUUID,
            Path:     path,
        })
    }

    if err := c.service.BulkCreate(banners); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, banners)
}

// GET /admin/banners
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

// PUT /admin/banners/:bannerId
func (c *BannerController) Update(ctx *gin.Context) {
    var banner models.ClientBanner
    if err := ctx.ShouldBindJSON(&banner); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    bannerID, _ := strconv.ParseInt(ctx.Param("bannerId"), 10, 64)
    banner.ID = uint(bannerID) // ✅ convert ke uint

    if err := c.service.Update(&banner); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, banner)
}

// DELETE /admin/banners/:bannerId
func (c *BannerController) Delete(ctx *gin.Context) {
    bannerID, err := strconv.ParseInt(ctx.Param("bannerId"), 10, 64)
    if err != nil || bannerID <= 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid banner id"})
        return
    }

    // ambil clientID dari query atau body (tergantung desain)
    clientID := ctx.Query("client_id")
    if clientID == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "missing client_id"})
        return
    }

    if err := c.service.Delete(bannerID, clientID); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "banner deleted"})
}
