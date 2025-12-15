package tenant

import (
    "net/http"

    // "backend_masjid/internal/models"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
)

type MasterClientController struct {
    service *services.MasterClientService
}

func NewMasterClientController(service *services.MasterClientService) *MasterClientController {
    return &MasterClientController{service}
}

// Tenant hanya bisa lihat konfigurasi masjidnya sendiri
func (c *MasterClientController) Get(ctx *gin.Context) {
    clientID := ctx.GetString("client_id")
    if clientID == "" {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing client_id claim"})
        return
    }

    client, err := c.service.GetByID(clientID)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    baseURL := "http://" + ctx.Request.Host
    if client.Logo != nil {
        client.LogoURL = services.GetFileURL(baseURL, *client.Logo)
    }
    if client.ConfigBackground != nil {
        client.BackgroundURL = services.GetFileURL(baseURL, *client.ConfigBackground)
    }
    if client.ConfigSoundAlert != nil {
        client.SoundURL = services.GetFileURL(baseURL, *client.ConfigSoundAlert)
    }

    ctx.JSON(http.StatusOK, client)
}


// Tenant hanya bisa update konfigurasi masjid miliknya
func (c *MasterClientController) Update(ctx *gin.Context) {
    clientID := ctx.GetString("client_id")
    if clientID == "" {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing client_id claim"})
        return
    }

    client, err := c.service.GetByID(clientID)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    // Ambil file logo
    if logoFile, _ := ctx.FormFile("logo"); logoFile != nil {
        path, err := services.SaveFile(ctx, logoFile)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        client.Logo = &path
    }

    // Ambil file background
    if bgFile, _ := ctx.FormFile("config_background"); bgFile != nil {
        path, err := services.SaveFile(ctx, bgFile)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        client.ConfigBackground = &path
    }

    // Ambil file sound alert
    if soundFile, _ := ctx.FormFile("config_sound_alert"); soundFile != nil {
        path, err := services.SaveFile(ctx, soundFile)
        if err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        client.ConfigSoundAlert = &path
    }

    // Ambil text field manual dari form-data
    if name := ctx.PostForm("name"); name != "" {
        client.Name = &name
    }
    if loc := ctx.PostForm("location"); loc != "" {
        client.Location = &loc
    }
    if tz := ctx.PostForm("timezone"); tz != "" {
        client.Timezone = &tz
    }
    if title := ctx.PostForm("config_title"); title != "" {
        client.ConfigTitle = &title
    }
    if rt := ctx.PostForm("running_text"); rt != "" {
        client.RunningText = &rt
    }

    // Boolean fields
    if hadis := ctx.PostForm("enable_hadis"); hadis != "" {
        val := hadis == "true"
        client.EnableHadis = &val
    }
    if hb := ctx.PostForm("enable_hari_besar"); hb != "" {
        val := hb == "true"
        client.EnableHariBesar = &val
    }
    if kal := ctx.PostForm("enable_kalender"); kal != "" {
        val := kal == "true"
        client.EnableKalender = &val
    }

    if err := c.service.Update(client); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "message": "config updated",
        "client":  client,
    })
}