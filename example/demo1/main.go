package main

import (
	fiber "github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
	_ "github.com/pagongamedev/go-dd/example/demo1/docs"
	"github.com/pagongamedev/go-dd/example/demo1/domain/demo"
)

func main() {
	portal := godd.NewPortal()
	appMain := appMain()

	portal.AppendApp(appMain, ":8083")
	portal.AppendApp(godd.AppGofiberAPIDocument(), ":8081")
	portal.AppendApp(godd.AppGofiberMetricsPrometheus(appMain), ":8082")

	portal.StartServer()
}

func appMain() *fiber.App {
	app := fiber.New()
	demo.Router(app, "/demo/v1")
	return app
}
