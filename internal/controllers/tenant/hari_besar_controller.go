package tenant

import (
    "net/http"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
)

type HariBesarController struct {
    service *services.HariBesarService
}

func NewHariBesarController(service *services.HariBesarService) *HariBesarController {
    return &HariBesarController{service}
}

func (c *HariBesarController) GetAll(ctx *gin.Context) {
    hariBesars, err := c.service.GetAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, hariBesars)
}

func (c *HariBesarController) GetByHoliday(ctx *gin.Context) {
    name := ctx.Param("name")
    hariBesars, err := c.service.GetByHoliday(name)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, hariBesars)
}