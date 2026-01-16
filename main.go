package main

import (
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/jejevj/go-aiocap/cmd"
	"github.com/jejevj/go-aiocap/config"
	"github.com/jejevj/go-aiocap/controller"
	_ "github.com/jejevj/go-aiocap/docs"
	"github.com/jejevj/go-aiocap/middleware"
	"github.com/jejevj/go-aiocap/repository"
	"github.com/jejevj/go-aiocap/routes"
	"github.com/jejevj/go-aiocap/service"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		cmd.Commands(db)
		return
	}

	var (
		jwtService service.JWTService = service.NewJWTService()

		userRepository repository.UserRepository = repository.NewUserRepository(db)
		userService    service.UserService       = service.NewUserService(userRepository, jwtService)
		userController controller.UserController = controller.NewUserController(userService)

		courseCustomerRepository repository.CourseCustomerRepository = repository.NewCourseCustomerRepository(db)
		courseCustomerService    service.CourseCustomerService       = service.NewCourseCustomerService(courseCustomerRepository, jwtService)
		courseCustomerController controller.CourseCustomerController = controller.NewCourseCustomerController(courseCustomerService)

		courseRepository repository.CourseRepository = repository.NewCourseRepository(db)
		courseService    service.CourseService       = service.NewCourseService(courseRepository, courseCustomerRepository, jwtService)
		courseController controller.CourseController = controller.NewCourseController(courseService)
	)

	server := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	server.Use(logger.New(logger.Config{
		CustomTags: map[string]logger.LogFunc{
			"custom_tag": func(output logger.Buffer, c *fiber.Ctx, data *logger.Data, extraParam string) (int, error) {
				return output.WriteString("it is a custom tag")
			},
		},
	}))

	file, err := os.OpenFile("./endpoint.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	server.Use(logger.New(logger.Config{
		Output: file,
	}))

	server.Use(middleware.CORSMiddleware())
	apiGroup := server.Group("/api")

	routes.User(apiGroup, userController, jwtService)
	routes.CourseCustomer(apiGroup, courseCustomerController, jwtService)
	routes.Course(apiGroup, courseController, jwtService)

	server.Get("/swagger/*", fiberSwagger.WrapHandler)
	server.Static("/assets", "./assets")
	server.Get("/metrics", monitor.New())
	server.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Listen(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
