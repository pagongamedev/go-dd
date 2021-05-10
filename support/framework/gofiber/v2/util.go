package v2

import (
	"net/http"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
	"github.com/pagongamedev/go-dd/support/framework/gofiber/v2/middleware"
)

//=============== Gofiber ======================

// AppAPIDocument Func
func AppAPIDocument() godd.InterfaceApp {
	goddApp, app := NewApp()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", http.StatusMovedPermanently)
	})
	app.Get("/swagger/*", swagger.Handler)
	return goddApp
}

// AppMetricsPrometheus Func
func AppMetricsPrometheus(mainApp godd.InterfaceApp) godd.InterfaceApp {
	goddApp, _ := NewApp()
	mdwPrometheus := middleware.NewPrometheus("fiber", "http")
	mdwPrometheus.Register(mainApp.App().(*fiber.App))
	mdwPrometheus.SetupPath(goddApp.App().(*fiber.App))
	return goddApp
}
