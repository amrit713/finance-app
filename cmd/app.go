package cmd

import (
	"log"

	"github.com/amirt713/finance-app/config"
	"github.com/amirt713/finance-app/internal/models"
	"github.com/amirt713/finance-app/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func Application() *fiber.App {

	app := fiber.New(fiber.Config{
		AppName: "Finance-app API",
	})

	config.ConnectDB()
	extQuery := `CREATE EXTENSION IF NOT EXISTS pgcrypto;`
	if err := config.DB.Exec(extQuery).Error; err != nil {
		log.Fatal("failed to create pgcrypto extension:", err)
	}

	err := config.DB.AutoMigrate(&models.User{}, &models.Budget{}, &models.Category{}, &models.Transaction{}, &models.Goal{}, &models.Notification{}, &models.Account{})

	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	routes.SetupRoutes(app)

	return app

}
