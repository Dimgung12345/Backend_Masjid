package admin

import (
    "net/http"
    "backend_masjid/internal/models"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
)

type MasterClientController struct {
    service *services.MasterClientService
}

func NewMasterClientController(service *services.MasterClientService) *MasterClientController {
    return &MasterClientController{service}
}


func (c *MasterClientController) Create(ctx *gin.Context) {
    var client models.MasterClient

    // ✅ pakai ShouldBindJSON untuk payload raw JSON
    if err := ctx.ShouldBindJSON(&client); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // kalau ada file upload (opsional)
    logoFile, _ := ctx.FormFile("logo")
    if logoFile != nil {
        path, err := services.SaveFile(ctx, logoFile)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        client.Logo = &path
    }

    bgFile, _ := ctx.FormFile("config_background")
    if bgFile != nil {
        path, err := services.SaveFile(ctx, bgFile)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        client.ConfigBackground = &path
    }

    soundFile, _ := ctx.FormFile("config_sound_alert")
    if soundFile != nil {
        path, err := services.SaveFile(ctx, soundFile)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        client.ConfigSoundAlert = &path
    }

    if err := c.service.Create(&client); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, client)
}

func (c *MasterClientController) GetAll(ctx *gin.Context) {
    clients, err := c.service.GetAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, clients)
}

func (c *MasterClientController) GetByID(ctx *gin.Context) {
    id := ctx.Param("id")
    client, err := c.service.GetByID(id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, client)
}

func (c *MasterClientController) Update(ctx *gin.Context) {
    var client models.MasterClient
    if err := ctx.ShouldBindJSON(&client); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.Update(&client); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, client)
}

func (c *MasterClientController) Delete(ctx *gin.Context) {
    id := ctx.Param("id")
    if err := c.service.Delete(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}