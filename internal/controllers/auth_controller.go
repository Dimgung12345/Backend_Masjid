package controllers

import (
    "net/http"
    "backend_masjid/internal/services"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "backend_masjid/internal/config"
)

type AuthController struct {
    service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
    return &AuthController{service}
}

// Login Admin
func (c *AuthController) LoginAdmin(ctx *gin.Context) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := c.service.LoginAdmin(req.Username, req.Password)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// Login DKM
// Login DKM + Display
func (c *AuthController) Login(ctx *gin.Context) {
    var req struct {
        Username string `json:"username"`
        Password string `json:"password"`
        Device   string `json:"device"`
    }
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := c.service.LoginDKM(req.Username, req.Password, req.Device)
    if err != nil {
        ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // Decode token untuk mengambil role
    parsed, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
        return []byte(config.Cfg.JwtSecret), nil
    })

    claims := parsed.Claims.(jwt.MapClaims)
    role := claims["role"]

    ctx.JSON(http.StatusOK, gin.H{
        "token": token,
        "role":  role,
    })
}