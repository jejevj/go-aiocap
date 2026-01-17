package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/go-aiocap/controller"
	"github.com/jejevj/go-aiocap/middleware"
	"github.com/jejevj/go-aiocap/service"
)

func Log(route fiber.Router, logController controller.LogController, jwtService service.JWTService) {
	routes := route.Group("/logs")

	routes.Get("", middleware.Authenticate(jwtService), logController.GetLogs)
	routes.Get("/:id", middleware.Authenticate(jwtService), logController.GetLogByID)
}
