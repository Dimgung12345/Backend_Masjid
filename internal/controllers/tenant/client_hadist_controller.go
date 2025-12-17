package tenant

import (
    "net/http"
    "strconv"

    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
)

type ClientHadistController struct {
    service *services.ClientHadistService
}

func NewClientHadistController(service *services.ClientHadistService) *ClientHadistController {
    return &ClientHadistController{service}
}

// Tenant lihat semua hadist global + status enabled/disabled
func (c *ClientHadistController) GetAll(ctx *gin.Context) {
    clientID := ctx.GetString("client_id")
    if clientID == "" {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing client_id claim"})
        return
    }

    limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
    offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

    hadists, err := c.service.GetAllByClient(clientID, limit, offset)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, hadists)
}

// Tenant disable hadist tertentu
func (c *ClientHadistController) Disable(ctx *gin.Context) {
    clientID := ctx.GetString("client_id")
    if clientID == "" {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing client_id claim"})
        return
    }

    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil || id <= 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid hadist id"})
        return
    }

    if err := c.service.Disable(clientID, uint(id)); err != nil {
        if err.Error() == "invalid hadist_id" {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err.Error() == "hadist not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "hadist disabled"})
}

// Tenant enable hadist tertentu
func (c *ClientHadistController) Enable(ctx *gin.Context) {
    clientID := ctx.GetString("client_id")
    if clientID == "" {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing client_id claim"})
        return
    }

    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil || id <= 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid hadist id"})
        return
    }

    if err := c.service.Enable(clientID, uint(id)); err != nil {
        if err.Error() == "invalid hadist_id" {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        if err.Error() == "hadist not found" {
            ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "hadist enabled"})
}

func (c *ClientHadistController) Search(ctx *gin.Context) {
    clientID := ctx.GetString("client_id")
    if clientID == "" {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing client_id claim"})
        return
    }

    keyword := ctx.Query("keyword")
    if keyword == "" {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "keyword is required"})
        return
    }

    limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
    offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

    hadists, err := c.service.SearchByKeyword(clientID, keyword, limit, offset)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, hadists)
}