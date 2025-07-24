package routes

import (
	"github.com/amirt713/finance-app/config"
	"github.com/amirt713/finance-app/internal/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	userRepo := repositories.NewUserRepository(config.DB)

	userRoute(api, userRepo)
	accountRoute(api, userRepo)
	budgetRoute(api, userRepo)
	categoryRoute(api, userRepo)
	transactionRoute(api, userRepo)
}
