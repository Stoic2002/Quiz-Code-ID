package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/codeid/hr-api/api/routes"
	configs "github.com/codeid/hr-api/internal/config"
	"github.com/codeid/hr-api/internal/models"
	"github.com/codeid/hr-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {

	//1. set environment (bisa cmd atau system environment)
	os.Setenv("APP_ENV", "development")

	//2.Load configuration
	config := configs.Load()

	//1. current code lebih minimalis
	db, err := database.InitDB(config)
	if err != nil {
		log.Fatal("failed to initialize database:%w", err)
	}

	// close db sebelum main application exit
	defer database.CloseDB(db)

	// Run auto migration, khusus bikin create table di database based on model
	// 1. bikin model dulu di go, lalu execute agar bisa create table di db
	if err := database.AutoMigrate(db, &models.Region{}, &models.Country{}); err != nil {
		log.Printf("Warning: Auto migration failed: %v", err)
	}

	// Set Gin mode based on environment
	if config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Setup routes, up running gin web server
	router := gin.Default()
	routes.SetupRoutes(router, db.DB)

	// Start server
	log.Printf("Server starting on %s in %s mode", config.Server.Address, config.Environment)

	// goroutine baru
	go func() {
		if err := router.Run(config.Server.Address); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 1. Graceful shutdown : nunggu operasi selesai baru shutdown server
	// 2. Tanpa Graceful Shutdowon : close connection,
	// tapi ada kemungkinan operasi seperti query masih jalan
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	/* //1. set datasourcename db config
	db, err := database.SetupDB() //posggre
	if err != nil {
		log.Fatal("failed to connect db:%w", err)
	}

	//1.1 initautomigrate --> baru
	database.InitAutoMigrate(db)

	//init repositories
	regionRepo := repositories.NewRegionRepository(db)

	//init service
	regionService := services.NewRegionService(regionRepo)

	//init handler/controler
	regionHandler := handlers.NewRegionHandler(regionService)

	//3. setup route
	router := gin.Default()

	//4. create router endpoint
	api := router.Group("/api")
	{
		//create region route
		regions := api.Group("/regions")
		{
			regions.GET("", regionHandler.GetRegions)
			regions.GET("/:id", regionHandler.GetRegion)
			regions.POST("", regionHandler.CreateRegion)
			regions.PUT("/:id", regionHandler.UpdateRegion)
			regions.DELETE("/:id", regionHandler.DeleteRegion)
			regions.GET("/countries", regionHandler.GetRegionsWithCountry)
			//add new endpoint
			regions.GET("/:id/countries", regionHandler.GetRegionByIdWithCountry)
		}

		//countries

	}

	//3.1 call handler
	router.GET("/", helloWorldHandler)

	//4. run webserver di port 8080
	router.Run(":8080") */
}
