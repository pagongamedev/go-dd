package main

import (
	fiber "github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
	"github.com/pagongamedev/go-dd/example/demo1/domain/demo"
)

func main() {
	portal := godd.NewPortal()
	appMain := appMain()

	portal.AppendApp(appMain, ":8080")
	portal.AppendApp(godd.AppAPIDocument(), ":8081")
	portal.AppendApp(godd.AppMetricsPrometheus(appMain), ":8082")

	portal.StartServer()
}

func appMain() *fiber.App {
	app := fiber.New()
	demo.Router(app, "/demo/v1")
	return app
}
