package admin

import (
    "encoding/json"
    "net/http"
    "strconv"
    "backend_masjid/internal/models"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
)

type KalenderController struct {
    service *services.KalenderService
}

func NewKalenderController(service *services.KalenderService) *KalenderController {
    return &KalenderController{service}
}
func (c *KalenderController) Create(ctx *gin.Context) {
    // Coba decode dulu sebagai array
    var kalenders []models.Kalender
    if err := ctx.ShouldBindJSON(&kalenders); err == nil {
        for _, k := range kalenders {
            if err := c.service.Create(&k); err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
        }
        ctx.JSON(http.StatusCreated, gin.H{
            "message": "Bulk kalenders created",
            "count":   len(kalenders),
        })
        return
    }

    // Kalau gagal parse array, coba parse single object
    var kalender models.Kalender
    if err := ctx.ShouldBindJSON(&kalender); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.Create(&kalender); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusCreated, kalender)
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

func (c *KalenderController) Update(ctx *gin.Context) {
    var kalender models.Kalender
    if err := ctx.ShouldBindJSON(&kalender); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.Update(&kalender); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, kalender)
}

func (c *KalenderController) Delete(ctx *gin.Context) {
    idStr := ctx.Param("id")
    id, _ := strconv.ParseInt(idStr, 10, 64)
    if err := c.service.Delete(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (c *KalenderController) BulkImport(ctx *gin.Context) {
    decoder := json.NewDecoder(ctx.Request.Body)

    // Pastikan body berupa array JSON
    token, err := decoder.Token()
    if err != nil || token != json.Delim('[') {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "expected JSON array"})
        return
    }

    batch := make([]models.Kalender, 0, 1000) // buffer batch
    total := 0

    for decoder.More() {
        var k models.Kalender
        if err := decoder.Decode(&k); err != nil {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        batch = append(batch, k)
        if len(batch) == 1000 {
            if err := c.service.BulkCreate(batch, 1000); err != nil {
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            total += len(batch)
            batch = batch[:0] // reset batch
        }
    }

    // Insert sisa batch
    if len(batch) > 0 {
        if err := c.service.BulkCreate(batch, 1000); err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        total += len(batch)
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "Bulk kalenders imported successfully",
        "count":   total,
    })
}