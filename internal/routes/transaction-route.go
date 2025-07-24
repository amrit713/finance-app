package routes

import (
	"github.com/amirt713/finance-app/config"
	"github.com/amirt713/finance-app/internal/controllers"
	"github.com/amirt713/finance-app/internal/middlewares"
	"github.com/amirt713/finance-app/internal/repositories"
	"github.com/amirt713/finance-app/internal/services"
	"github.com/gofiber/fiber/v2"
)

func transactionRoute(api fiber.Router, userRepo *repositories.UserRepository) {

	transactionRepo := repositories.NewTransactionRepository(config.DB)
	transactionService := services.NewTransactionService(transactionRepo)
	transactionController := controllers.NewTransactionController(transactionService)

	api.Get("/accounts/:accountId/transactions", middlewares.Protected(userRepo), transactionController.GetAccountTransactions)

	accountApi := api.Group("/transactions", middlewares.Protected(userRepo))

	accountApi.Get("", transactionController.GetAllTransactions)
	accountApi.Post("", transactionController.CreateTransaction)
	accountApi.Get("/:id", transactionController.GetTransaction)
	accountApi.Put("/:id", transactionController.UpdateTransaction)
	accountApi.Delete("/:id", transactionController.DeleteTransaction)

}
