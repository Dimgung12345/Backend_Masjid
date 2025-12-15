package admin

import (
    "net/http"
    "strconv"
    "backend_masjid/internal/models"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
)

type HadistController struct {
    service *services.HadistService
}

func NewHadistController(service *services.HadistService) *HadistController {
    return &HadistController{service}
}

func (c *HadistController) Create(ctx *gin.Context) {
    // Coba decode dulu sebagai array
    var hadists []models.Hadist
    if err := ctx.ShouldBindJSON(&hadists); err == nil {
        // Kalau sukses parse array
        for _, h := range hadists {
            if err := c.service.Create(&h); err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
        }
        ctx.JSON(http.StatusCreated, gin.H{"message": "Bulk hadists created", "count": len(hadists)})
        return
    }

    // Kalau gagal parse array, coba parse single object
    var hadist models.Hadist
    if err := ctx.ShouldBindJSON(&hadist); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.Create(&hadist); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, hadist)
}

func (c *HadistController) GetAll(ctx *gin.Context) {
    limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
    offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

    hadists, err := c.service.GetAll(limit, offset)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, hadists)
}

func (c *HadistController) GetByID(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, _ := strconv.ParseInt(idStr, 10, 64)
    hadist, err := c.service.GetByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, hadist)
}

func (c *HadistController) Update(ctx *gin.Context) {
    var hadist models.Hadist
    if err := ctx.ShouldBindJSON(&hadist); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.Update(&hadist); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, hadist)
}

func (c *HadistController) Delete(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, _ := strconv.ParseInt(idStr, 10, 64)
    if err := c.service.Delete(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}