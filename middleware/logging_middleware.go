package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/go-aiocap/service"
)

func LoggingMiddleware(logService service.LogService, jwtService service.JWTService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Capture request body BEFORE next
		requestBody := ctx.Body()

		// Try to extract user_id from JWT (optional auth)
		var userID string
		authHeader := ctx.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
			if uid, err := jwtService.GetUserIDByToken(tokenStr); err == nil {
				userID = uid
				ctx.Locals("user_id", uid) // optional, for controllers
			}
		}

		// Continue request chain
		err := ctx.Next()

		// If no authenticated user → skip logging
		if userID == "" {
			return err
		}

		// Capture response body AFTER next
		responseBody := ""
		if ctx.Response().Body() != nil {
			responseBody = string(ctx.Response().Body())
		}

		action := getActionDescription(ctx.Method(), ctx.Path())

		logService.LogAction(
			ctx.Context(),
			action,
			ctx.Path(),
			ctx.Method(),
			userID,
			string(requestBody),
			responseBody,
		)

		return err
	}
}

func getActionDescription(method, path string) string {
	switch method {
	case "POST":
		return "Menambahkan data"
	case "GET":
		return "Mengambil data"
	case "PUT", "PATCH":
		return "Memperbarui data"
	case "DELETE":
		return "Menghapus data"
	default:
		return "Melakukan aksi"
	}
}
