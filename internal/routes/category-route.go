package routes

import (
	"github.com/amirt713/finance-app/config"
	"github.com/amirt713/finance-app/internal/controllers"
	"github.com/amirt713/finance-app/internal/middlewares"
	"github.com/amirt713/finance-app/internal/repositories"
	"github.com/amirt713/finance-app/internal/services"
	"github.com/gofiber/fiber/v2"
)

func categoryRoute(api fiber.Router, userRepo *repositories.UserRepository) {

	categoryRepo := repositories.NewCategoryRepository(config.DB)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)

	accountApi := api.Group("/categories", middlewares.Protected(userRepo))

	accountApi.Get("", categoryController.GetAllCategories)
	accountApi.Post("", categoryController.CreateCategory)
	accountApi.Get("/:id", categoryController.GetCategory)
	accountApi.Put("/:id", categoryController.UpdateCategory)
	accountApi.Delete("/:id", categoryController.DeleteCategory)

}
