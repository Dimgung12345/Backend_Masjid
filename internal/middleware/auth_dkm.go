package middleware

import (
    "net/http"
    "strings"

    "backend_masjid/internal/config"
    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

func AuthRole(allowedRoles ...string) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        tokenString := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
        if tokenString == "" {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            ctx.Abort()
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte(config.Cfg.JwtSecret), nil
        })
        if err != nil || !token.Valid {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            ctx.Abort()
            return
        }

        claims := token.Claims.(jwt.MapClaims)
        role := claims["role"].(string)

        ok := false
        for _, r := range allowedRoles {
            if role == r {
                ok = true
                break
            }
        }
        if !ok {
            ctx.JSON(http.StatusForbidden, gin.H{"error": "forbidden"})
            ctx.Abort()
            return
        }

        if clientID, ok := claims["client_id"].(string); ok {
            ctx.Set("client_id", clientID)
        }
        ctx.Next()
    }
}