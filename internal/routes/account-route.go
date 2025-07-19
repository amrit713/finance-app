package routes

import (
	"github.com/amirt713/finance-app/config"
	"github.com/amirt713/finance-app/internal/controllers"
	"github.com/amirt713/finance-app/internal/middlewares"
	"github.com/amirt713/finance-app/internal/repositories"
	"github.com/amirt713/finance-app/internal/services"
	"github.com/gofiber/fiber/v2"
)

func accountRoute(api fiber.Router, userRepo *repositories.UserRepository) {

	accountRepo := repositories.NewAccountRepository(config.DB)
	accountService := services.NewAccountService(accountRepo)
	accountController := controllers.NewAccountController(accountService)

	accountApi := api.Group("/accounts", middlewares.Protected(userRepo))

	accountApi.Get("", accountController.GetAllAccounts)
	accountApi.Post("", accountController.CreateAccount)
	accountApi.Get("/:id", accountController.GetAccount)
	accountApi.Put("/:id", accountController.UpdateAccount)
	accountApi.Delete("/:id", accountController.DeleteAccount)

}
