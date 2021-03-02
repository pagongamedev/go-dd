package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/pagongamedev/go-dd/example/ex003_repo_switcher/notexample"
	"github.com/pagongamedev/go-dd/example/ex003_repo_switcher/repo01"
	"github.com/pagongamedev/go-dd/example/ex003_repo_switcher/repo02"
	"github.com/pagongamedev/go-dd/example/ex003_repo_switcher/service"

	godd "github.com/pagongamedev/go-dd"
	goddPortal "github.com/pagongamedev/go-dd/portal"
)

// Can test api by localhost:8081/hello/v1/hello

// Example ex003_repo_switcher
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

	//  Repository Pattern Switcher
	env := "localhost"
	funcRepository := godd.EnvironmentSwitcher(env, 0, 0, 1, 1, 1,
		repo01.NewRepository, repo02.NewRepository)

	repoHello, err := funcRepository.(func() (service.Repository, error))()
	godd.MustError(err)

	// Manage Service
	serviceHello, err := service.NewService(repoHello)
	godd.MustError(err)

	notexample.Router(app, "/hello/v1", serviceHello)
	return app
}
