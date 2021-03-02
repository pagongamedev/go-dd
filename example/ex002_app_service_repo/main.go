package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/pagongamedev/go-dd/example/ex002_app_service_repo/notexample"
	"github.com/pagongamedev/go-dd/example/ex002_app_service_repo/service"

	godd "github.com/pagongamedev/go-dd"
	goddPortal "github.com/pagongamedev/go-dd/portal"
)

// Can test api by localhost:8081/hello/v1/hello

// Example ex002_app_service_repo
func main() {
	portal := goddPortal.New()
	portal.AppendApp(appMain(), ":8081")
	portal.StartServer()
}

func appMain() *fiber.App {
	app := fiber.New()

	// Use App Middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// Manage Service And Repository Pattern
	repoHello, err := service.NewRepository()
	godd.MustError(err)
	serviceHello, err := service.NewService(repoHello)
	godd.MustError(err)

	notexample.Router(app, "/hello/v1", serviceHello)
	return app
}
