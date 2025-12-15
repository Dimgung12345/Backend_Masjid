package admin

import (
    "net/http"
    "strconv"
    "backend_masjid/internal/models"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
)

type HariBesarController struct {
    service *services.HariBesarService
}

func NewHariBesarController(service *services.HariBesarService) *HariBesarController {
    return &HariBesarController{service}
}

func (c *HariBesarController) Create(ctx *gin.Context) {
    // Coba parse array dulu
    var hariBesars []models.HariBesar
    if err := ctx.ShouldBindJSON(&hariBesars); err == nil {
        // Kalau sukses parse array
        for _, hb := range hariBesars {
            if err := c.service.Create(&hb); err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
        }
        ctx.JSON(http.StatusCreated, gin.H{
            "message": "Bulk hari besar created",
            "count":   len(hariBesars),
        })
        return
    }

    // Kalau bukan array, parse single object
    var hariBesar models.HariBesar
    if err := ctx.ShouldBindJSON(&hariBesar); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.Create(&hariBesar); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, hariBesar)
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

func (c *HariBesarController) Update(ctx *gin.Context) {
        var Haris []models.HariBesar
    if err := ctx.ShouldBindJSON(&Haris); err == nil {
        // Kalau sukses parse array
        for _, j := range Haris {
            if err := c.service.Create(&j); err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
        }
        ctx.JSON(http.StatusCreated, gin.H{"message": "Bulk hadists created", "count": len(Haris)})
        return
    }
        var hariBesar models.HariBesar
    if err := ctx.ShouldBindJSON(&hariBesar); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.Update(&hariBesar); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, hariBesar)
}

func (c *HariBesarController) Delete(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, _ := strconv.ParseInt(idStr, 10, 64)
    if err := c.service.Delete(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}