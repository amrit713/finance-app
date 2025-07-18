package routes

import (
	"github.com/amirt713/finance-app/config"
	"github.com/amirt713/finance-app/internal/controllers"
	"github.com/amirt713/finance-app/internal/middlewares"
	"github.com/amirt713/finance-app/internal/repositories"
	"github.com/amirt713/finance-app/internal/services"
	"github.com/gofiber/fiber/v2"
)

func userRoute(api fiber.Router) {

	userRepo := repositories.NewUserRepository(config.DB)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	authApi := api.Group("/auth")
	authApi.Post("/sign-up", authController.Register)
	authApi.Post("/login", authController.Login)
	authApi.Post("/logout", authController.Logout)

	userApi := api.Group("/users")
	userApi.Get("", middlewares.Protected(userRepo), userController.GetAllUsers)
	userApi.Put("/profile", middlewares.Protected(userRepo), userController.EditMe)

}
