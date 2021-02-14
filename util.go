package godd

import (
	"log"
	"net/http"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/pagongamedev/godd/middleware"
)

// =====================================================================
//                              Add On
// =====================================================================

// MustError Func
func MustError(err error, strList ...string) {
	if err != nil {
		if strList != nil {
			log.Fatal(strList)
		} else {
			log.Fatal("Error : ", err)
		}
	}
}

// AddAPIGetHealth Func
func addAPIGetHealth(app *fiber.App) {
	app.Get("/health", handlerHealth())
}

func handlerHealth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"success": true})
	}
}

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
