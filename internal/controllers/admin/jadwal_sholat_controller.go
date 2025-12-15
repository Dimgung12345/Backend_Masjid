package admin

import (
    "net/http"
    "strconv"
    "backend_masjid/internal/models"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
)

type JadwalController struct {
    service *services.JadwalService
}

func NewJadwalController(service *services.JadwalService) *JadwalController {
    return &JadwalController{service}
}

func (c *JadwalController) Create(ctx *gin.Context) {
    // Coba decode dulu sebagai array
    var jadwals []models.JadwalSholat
    if err := ctx.ShouldBindJSON(&jadwals); err == nil {
        // Kalau sukses parse array
        for _, j := range jadwals {
            if err := c.service.Create(&j); err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
        }
        ctx.JSON(http.StatusCreated, gin.H{"message": "Bulk hadists created", "count": len(jadwals)})
        return
    }

    // Kalau gagal parse array, coba parse single object
    var jadwal models.JadwalSholat
    if err := ctx.ShouldBindJSON(&jadwal); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.Create(&jadwal); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, jadwal)
}

func (c *JadwalController) GetAll(ctx *gin.Context) {
    jadwals, err := c.service.GetAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, jadwals)
}

func (c *JadwalController) GetByDate(ctx *gin.Context) {
    date := ctx.Param("date")
    jadwal, err := c.service.GetByDate(date)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, jadwal)
}

func (c *JadwalController) Update(ctx *gin.Context) {
    var jadwal models.JadwalSholat
    if err := ctx.ShouldBindJSON(&jadwal); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.Update(&jadwal); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, jadwal)
}

func (c *JadwalController) Delete(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, _ := strconv.ParseInt(idStr, 10, 64)
    if err := c.service.Delete(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}