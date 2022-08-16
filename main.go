package main

import (
	"fmt"
	"os"
	"sync"

	"js-ptpos/config"
	db "js-ptpos/config"
	"js-ptpos/controller"
	"js-ptpos/exception"
	"js-ptpos/repository"
	"js-ptpos/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := db.InitEcha()
		if err != nil {
			fmt.Println("Oracle Echa:", err.Error())
			os.Exit(1)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := db.InitOltp()
		if err != nil {
			fmt.Println("Oracle Oltp:", err.Error())
			os.Exit(1)
		}
	}()

	wg.Wait()
	fmt.Println("Echa  :", db.EngineEcha.DataSourceName())
	fmt.Println("Oltp  :", db.EngineOltp.DataSourceName())

	configuration := config.New()

	// Setup Repository
	claimRepository := repository.NewClaimRepository(&configuration)
	beasiswaRepository := repository.NewBeasiswaRepository()
	claimJPRepository := repository.NewClaimJPRepository()
	claimJHTRepository := repository.NewJHTClaimRepository()
	storageRepository := repository.NewStorageRepository(&configuration)
	emailRepository := repository.NewEmailRepository()

	// Setup Service
	claimService := service.NewClaimService(&claimRepository)
	beasiswaService := service.NewBeasiswaService(&beasiswaRepository)
	claimJPService := service.NewClaimJPService(&claimJPRepository)
	claimJHTService := service.NewJHTClaimService(&claimJHTRepository)
	storageService := service.NewStorageService(&storageRepository)
	emailService := service.NewEmailService(&emailRepository)

	// Setup Controller
	claimController := controller.NewClaimController(&claimService)
	beasiswaController := controller.NewBeasiswaController(&beasiswaService)
	claimJPController := controller.NewClaimJPController(&claimJPService)
	claimJHTController := controller.NewClaimJHTController(&claimJHTService)
	storageController := controller.NewStorageController(&storageService)
	emailController := controller.NewEmailController(&emailService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Logging
	app.Use(logger.New())

	// Setup Routing
	claimController.Route(app)
	beasiswaController.Route(app)
	claimJPController.Route(app)
	claimJHTController.Route(app)
	storageController.Route(app)
	emailController.Route(app)

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)
}
