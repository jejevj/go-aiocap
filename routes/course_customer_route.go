package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jejevj/go-aiocap/controller"
	"github.com/jejevj/go-aiocap/service"
)

func CourseCustomer(route fiber.Router, courseCustomerController controller.CourseCustomerController, jwtService service.JWTService) {
	routes := route.Group("/course-customer")

	routes.Post("", courseCustomerController.AddCourseCustomer)
}
