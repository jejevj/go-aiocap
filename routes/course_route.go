package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/go-aiocap/controller"
	"github.com/jejevj/go-aiocap/service"
)

func Course(route fiber.Router, courseController controller.CourseController, jwtService service.JWTService) {
	routes := route.Group("/course")

	routes.Post("", courseController.AddCourse)
}
