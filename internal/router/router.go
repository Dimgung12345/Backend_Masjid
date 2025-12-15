package router

import (
    "backend_masjid/internal/controllers"
    "backend_masjid/internal/controllers/admin"
    "backend_masjid/internal/controllers/tenant"
    "backend_masjid/internal/middleware"
    "backend_masjid/internal/services"
    "github.com/gin-contrib/cors"
    "time"
    "github.com/gin-gonic/gin"
)

func SetupRouter(
    hadistService *services.HadistService,
    jadwalService *services.JadwalService,
    kalenderService *services.KalenderService,
    hariBesarService *services.HariBesarService,
    masterClientService *services.MasterClientService,
    dkmUserService *services.DkmUserService,
    authService *services.AuthService,
    bannerService *services.BannerService,
    clientHadistService *services.ClientHadistService,
) *gin.Engine {
    r := gin.Default()

        // ✅ Pasang CORS global di sini
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // sesuaikan origin FE kamu
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // Auth routes
    authCtrl := controllers.NewAuthController(authService)
    r.POST("/login/admin", authCtrl.LoginAdmin)
    r.POST("/login", authCtrl.Login)

    // Admin routes
    adminGroup := r.Group("/admin", middleware.AuthAdmin())
    {
        hadistCtrl := admin.NewHadistController(hadistService)
        adminGroup.POST("/hadist", hadistCtrl.Create)
        adminGroup.GET("/hadist", hadistCtrl.GetAll)
        adminGroup.GET("/hadist/:id", hadistCtrl.GetByID)
        adminGroup.PUT("/hadist", hadistCtrl.Update)
        adminGroup.DELETE("/hadist/:id", hadistCtrl.Delete)

        jadwalCtrl := admin.NewJadwalController(jadwalService)
        adminGroup.POST("/jadwal", jadwalCtrl.Create)
        adminGroup.GET("/jadwal", jadwalCtrl.GetAll)
        adminGroup.GET("/jadwal/:date", jadwalCtrl.GetByDate)
        adminGroup.PUT("/jadwal", jadwalCtrl.Update)
        adminGroup.DELETE("/jadwal/:id", jadwalCtrl.Delete)

        kalenderCtrl := admin.NewKalenderController(kalenderService)
        adminGroup.POST("/kalender", kalenderCtrl.Create)
        adminGroup.POST("/kalender/bulk", kalenderCtrl.BulkImport)
        adminGroup.GET("/kalender", kalenderCtrl.GetAll)
        adminGroup.GET("/kalender/:date", kalenderCtrl.GetByDate)
        adminGroup.PUT("/kalender", kalenderCtrl.Update)
        adminGroup.DELETE("/kalender/:id", kalenderCtrl.Delete)

        hariBesarCtrl := admin.NewHariBesarController(hariBesarService)
        adminGroup.POST("/hari-besar", hariBesarCtrl.Create)
        adminGroup.GET("/hari-besar", hariBesarCtrl.GetAll)
        adminGroup.GET("/hari-besar/:name", hariBesarCtrl.GetByHoliday)
        adminGroup.PUT("/hari-besar", hariBesarCtrl.Update)
        adminGroup.DELETE("/hari-besar/:id", hariBesarCtrl.Delete)

        clientCtrl := admin.NewMasterClientController(masterClientService)
        adminGroup.POST("/client", clientCtrl.Create)
        adminGroup.GET("/client", clientCtrl.GetAll)
        adminGroup.GET("/client/:id", clientCtrl.GetByID)
        adminGroup.PUT("/client", clientCtrl.Update)
        adminGroup.DELETE("/client/:id", clientCtrl.Delete)

        dkmCtrl := admin.NewDkmUserController(dkmUserService)
        adminGroup.POST("/dkm", dkmCtrl.Create)
        adminGroup.GET("/dkm", dkmCtrl.GetAll)
        adminGroup.GET("/dkm/:username", dkmCtrl.GetByUsername)
        adminGroup.PUT("/dkm", dkmCtrl.Update)
        adminGroup.DELETE("/dkm/:id", dkmCtrl.Delete)

        adminBanner := admin.NewBannerController(bannerService)
        adminGroup.POST("/banners", adminBanner.Create)
        adminGroup.GET("/banners", adminBanner.GetAll)
        adminGroup.PUT("/banners/:bannerId", adminBanner.Update)
        adminGroup.DELETE("/banners/:bannerId", adminBanner.Delete)

    }

    // Tenant routes
    tenantGroup := r.Group("/tenant", middleware.AuthRole("dkm", "display"))
    {
        tenantHadist := tenant.NewHadistController(hadistService)
        tenantGroup.GET("/hadist", tenantHadist.GetAll)
        tenantGroup.GET("/hadist/:id", tenantHadist.GetByID)

        tenantClientHadist := tenant.NewClientHadistController(clientHadistService)
        tenantGroup.GET("/hadists", tenantClientHadist.GetAll)
        tenantGroup.PUT("/hadists/:id/disable", tenantClientHadist.Disable)
        tenantGroup.PUT("/hadists/:id/enable", tenantClientHadist.Enable)

        tenantJadwal := tenant.NewJadwalController(jadwalService)
        tenantGroup.GET("/jadwal", tenantJadwal.GetAll)
        tenantGroup.GET("/jadwal/date/:date", tenantJadwal.GetByDate)
        tenantGroup.GET("/jadwal/lokasi/:lokasi", tenantJadwal.GetByLokasi)

        tenantKalender := tenant.NewKalenderController(kalenderService)
        tenantGroup.GET("/kalender", tenantKalender.GetAll)
        tenantGroup.GET("/kalender/:date", tenantKalender.GetByDate)

        tenantHariBesar := tenant.NewHariBesarController(hariBesarService)
        tenantGroup.GET("/hari-besar", tenantHariBesar.GetAll)
        tenantGroup.GET("/hari-besar/:name", tenantHariBesar.GetByHoliday)

        tenantClient := tenant.NewMasterClientController(masterClientService)
        tenantGroup.GET("/client", tenantClient.Get)
        tenantGroup.PUT("/client", tenantClient.Update)

        tenantBanner := tenant.NewBannerController(bannerService)
        tenantGroup.GET("/banners", tenantBanner.GetAll)
        tenantGroup.PUT("/banners/:bannerId", tenantBanner.Update)
    }

    return r
}