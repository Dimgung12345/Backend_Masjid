package admin

import (
    "net/http"
    "backend_masjid/internal/models"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"

)

type DkmUserController struct {
    service *services.DkmUserService
}

func NewDkmUserController(service *services.DkmUserService) *DkmUserController {
    return &DkmUserController{service}
}

func (c *DkmUserController) Create(ctx *gin.Context) {
    var user models.DkmUser
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Hash password sebelum simpan
    hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
        return
    }
    user.Password = string(hashed)

    if err := c.service.Create(&user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{
        "message": "DKM user created",
        "user":    user,
    })
}


func (c *DkmUserController) GetAll(ctx *gin.Context) {
    users, err := c.service.GetAll()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, users)
}

func (c *DkmUserController) GetByUsername(ctx *gin.Context) {
    username := ctx.Param("username")
    user, err := c.service.GetByUsername(username)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, user)
}

func (c *DkmUserController) Update(ctx *gin.Context) {
    var user models.DkmUser
    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if err := c.service.Update(&user); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, user)
}

func (c *DkmUserController) Delete(ctx *gin.Context) {
    id := ctx.Param("id")
    if err := c.service.Delete(id); err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}