package v2

import (
	"net/http"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
	"github.com/pagongamedev/go-dd/support/gofiber/v2/middleware"
)

//=============== Gofiber ======================

// AppAPIDocument Func
func AppAPIDocument() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html", http.StatusMovedPermanently)
	})
	app.Get("/swagger/*", swagger.Handler)
	return app
}

// AppMetricsPrometheus Func
func AppMetricsPrometheus(mainApp *fiber.App) *fiber.App {
	app := fiber.New()
	mdwPrometheus := middleware.NewPrometheus("fiber", "http")
	mdwPrometheus.Register(mainApp)
	mdwPrometheus.SetupPath(app)
	return app
}

func NewApp() (*godd.App, *fiber.App) {
	app := fiber.New()
	goddApp := &godd.App{App: app, FrameWork: godd.FrameWorkGofiberV2}

	return goddApp, app
}
