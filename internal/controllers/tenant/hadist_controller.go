package tenant

import (
    "net/http"
    "strconv"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
)

type HadistController struct {
    service *services.HadistService
}

func NewHadistController(service *services.HadistService) *HadistController {
    return &HadistController{service}
}

func (c *HadistController) GetAll(ctx *gin.Context) {
    limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
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