package main

import (
    "time"

    "github.com/gin-contrib/cors"
    // "github.com/gin-gonic/gin"

    "backend_masjid/internal/config"
    "backend_masjid/internal/db"
    "backend_masjid/internal/router"
    "backend_masjid/internal/repository"
    "backend_masjid/internal/services"
)

func main() {
    // 1. Load config
    config.LoadConfig()

    // 2. Init DB
    db.InitDB()

    // 3. Init repositories
    hadistRepo := repository.NewHadistRepository(db.DB)
    clientHadistRepo := repository.NewClientHadistRepository(db.DB)
    jadwalRepo := repository.NewJadwalRepository(db.DB)
    kalenderRepo := repository.NewKalenderRepository(db.DB)
    hariBesarRepo := repository.NewHariBesarRepository(db.DB)
    masterClientRepo := repository.NewMasterClientRepository(db.DB)
    dkmUserRepo := repository.NewDkmUserRepository(db.DB)
    adminRepo := repository.NewAdminUserRepository(db.DB)
    dkmRepo := repository.NewDkmUserRepository(db.DB)
    bannerRepo := repository.NewBannerRepository(db.DB)

    // 4. Init services
    hadistService := services.NewHadistService(hadistRepo)
    clientHadistService := services.NewClientHadistService(clientHadistRepo, hadistRepo)
    jadwalService := services.NewJadwalService(jadwalRepo)
    kalenderService := services.NewKalenderService(kalenderRepo)
    hariBesarService := services.NewHariBesarService(hariBesarRepo)
    masterClientService := services.NewMasterClientService(masterClientRepo)
    dkmUserService := services.NewDkmUserService(dkmUserRepo)
    authService := services.NewAuthService(adminRepo, dkmRepo)
    bannerService := services.NewBannerService(bannerRepo)

    // 5. Setup router
    r := router.SetupRouter(
        hadistService,
        jadwalService,
        kalenderService,
        hariBesarService,
        masterClientService,
        dkmUserService,
        authService,
        bannerService,
        clientHadistService,
    )

    // 👉 Pasang CORS dari gin-contrib
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))


    r.Static("/storage", "./storage")

    
    // 6. Run server
    r.Run(":" + config.Cfg.AppPort)
}