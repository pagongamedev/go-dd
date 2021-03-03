package main

import (
	"log"

	fiber "github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
	_ "github.com/pagongamedev/go-dd/example/old_demo1/docs"
	"github.com/pagongamedev/go-dd/example/old_demo1/domain/demo"
	goddPortal "github.com/pagongamedev/go-dd/portal"
	goddGofiber "github.com/pagongamedev/go-dd/support/gofiber"
)

type a struct {
}

func (a *a) Close() error {
	log.Println("Close")
	return nil
}

func main() {
	portal := goddPortal.New()
	appMain := appMain()

	aa := a{}

	portal.AppendApp(appMain, ":8081")
	portal.AppendApp(goddGofiber.AppAPIDocument(), ":8082")
	portal.AppendApp(goddGofiber.AppMetricsPrometheus(appMain), ":8083")

	portal.AppendDeferClose(godd.DeferClose{Name: "A", I: &aa})

	portal.StartServer()
}

func appMain() *fiber.App {
	app := fiber.New()
	demo.Router(app, "/demo/v1")
	return app
}
