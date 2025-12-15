package tenant

import (
    "net/http"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
)

type KalenderController struct {
    service *services.KalenderService
}

func NewKalenderController(service *services.KalenderService) *KalenderController {
    return &KalenderController{service}
}

func (c *KalenderController) GetAll(ctx *gin.Context) {
    kalenders, err := c.service.GetAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, kalenders)
}

func (c *KalenderController) GetByDate(ctx *gin.Context) {
    date := ctx.Param("date")
    kalender, err := c.service.GetByDate(date)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, kalender)
}