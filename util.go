package godd

import (
	"log"
	"net/http"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/pagongamedev/go-dd/middleware"
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
func addAPIGetHealth(app InterfaceApp) {
	app.Get("/health", handlerHealth())
}

func handlerHealth() Handler {
	return func(ctx InterfaceContext) error {
		return ctx.Response(Map{"success": true})
	}
}

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
