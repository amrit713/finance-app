package routes

import (
	"github.com/amirt713/finance-app/config"
	"github.com/amirt713/finance-app/internal/controllers"
	"github.com/amirt713/finance-app/internal/middlewares"
	"github.com/amirt713/finance-app/internal/repositories"
	"github.com/amirt713/finance-app/internal/services"
	"github.com/gofiber/fiber/v2"
)

func budgetRoute(api fiber.Router, userRepo *repositories.UserRepository) {

	budgetRepo := repositories.NewBudgetRepository(config.DB)
	budgetService := services.NewBudgetService(budgetRepo)
	budgetController := controllers.NewBudgetController(budgetService)

	budgetApi := api.Group("/budgets", middlewares.Protected(userRepo))

	budgetApi.Post("", budgetController.CreateBudget)
	budgetApi.Get("/:id", budgetController.GetBudget)
	budgetApi.Delete("/:id", budgetController.DeleteBudget)
	budgetApi.Patch("/:id", budgetController.UpdateBudget)

}
