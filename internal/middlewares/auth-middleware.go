package middlewares

import (
	"strings"

	"github.com/amirt713/finance-app/internal/repositories"
	"github.com/amirt713/finance-app/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected(userRepo *repositories.UserRepository) fiber.Handler {

	return func(ctx *fiber.Ctx) error {

		var jwtToken string

		//1. check Auth Header
		authHeader := ctx.Get("Authorization")
		if authHeader != "" {
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
				jwtToken = parts[1]
			}
		}

		// 2. Fallback: Check Cookie if no jwtToken from header

		if jwtToken == "" {
			cookieToken := ctx.Cookies("jwt_token")
			if cookieToken != "" {
				jwtToken = cookieToken
			}
		}

		// 3. If still no jwtToken, return unauthorized
		if jwtToken == "" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authentication token",
			})
		}

		token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (any, error) {
			return utils.JWTSecret, nil
		})

		if err != nil || !token.Valid {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"statusCode": fiber.StatusUnauthorized,
				"error":      "unauthorized",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Failed to parse claims",
			})
		}

		userID, ok := claims["user_id"].(string)

		if !ok {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid user ID in token"})
		}

		user, err := userRepo.FindByID(userID)

		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
		}

		ctx.Locals("user", user)
		return ctx.Next()
	}

}
