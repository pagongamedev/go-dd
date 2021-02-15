package godd

import (
	"net/http"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/pagongamedev/go-dd/middleware"
)

//=============== Gofiber ======================

// AppGofiberAPIDocument Func
func AppGofiberAPIDocument() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", http.StatusMovedPermanently)
	})
	app.Get("/swagger/*", swagger.Handler)
	return app
}

// AppGofiberMetricsPrometheus Func
func AppGofiberMetricsPrometheus(mainApp *fiber.App) *fiber.App {
	app := fiber.New()
	mdwPrometheus := middleware.NewPrometheus("fiber", "http")
	mdwPrometheus.Register(mainApp)
	mdwPrometheus.SetupPath(app)
	return app
}
