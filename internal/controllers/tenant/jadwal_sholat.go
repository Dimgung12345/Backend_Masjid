package tenant

import (
    "net/http"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
)

type JadwalController struct {
    service *services.JadwalService
}

func NewJadwalController(service *services.JadwalService) *JadwalController {
    return &JadwalController{service}
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

func (c *JadwalController) GetByLokasi(ctx *gin.Context) {
    lokasi := ctx.Param("lokasi")
    jadlok, err := c.service.GetByLokasi(lokasi)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, jadlok)
}