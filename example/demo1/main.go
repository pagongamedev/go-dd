package main

import (
	fiber "github.com/gofiber/fiber/v2"
	_ "github.com/pagongamedev/go-dd/example/demo1/docs"
	"github.com/pagongamedev/go-dd/example/demo1/domain/demo"
	goddPortal "github.com/pagongamedev/go-dd/portal"
	goddGofiber "github.com/pagongamedev/go-dd/support/gofiber"
)

func main() {
	portal := goddPortal.New()
	appMain := appMain()

	portal.AppendApp(appMain, ":8083")
	portal.AppendApp(goddGofiber.AppAPIDocument(), ":8081")
	portal.AppendApp(goddGofiber.AppMetricsPrometheus(appMain), ":8082")

	portal.StartServer()
}

func appMain() *fiber.App {
	app := fiber.New()
	demo.Router(app, "/demo/v1")
	return app
}
