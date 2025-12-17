package main

import (
    "time"

    "github.com/gin-contrib/cors"

    "backend_masjid/internal/config"
    "backend_masjid/internal/db"
    "backend_masjid/internal/router"
    "backend_masjid/internal/repository"
    "backend_masjid/internal/services"
    "backend_masjid/internal/cache"
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

    // 4. Init cache manager (misalnya TTL 24 jam)
    cacheManager := cache.NewCacheManager(24 * time.Hour)

    // 5. Init services dengan cache
    hadistService := services.NewHadistService(hadistRepo)
    clientHadistService := services.NewClientHadistService(clientHadistRepo, hadistRepo, cacheManager)
    jadwalService := services.NewJadwalService(jadwalRepo)
    kalenderService := services.NewKalenderService(kalenderRepo)
    hariBesarService := services.NewHariBesarService(hariBesarRepo)
    masterClientService := services.NewMasterClientService(masterClientRepo, cacheManager)
    dkmUserService := services.NewDkmUserService(dkmUserRepo)
    authService := services.NewAuthService(adminRepo, dkmRepo)
    bannerService := services.NewBannerService(bannerRepo, cacheManager)

    // 6. Setup router
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

    // 7. Run server
    r.Run(":" + config.Cfg.AppPort)
}