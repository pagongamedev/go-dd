package demo

import (
	"github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
)

var ms *godd.MicroService

// Router Func
func Router(app *fiber.App, path string) *godd.MicroService {

	ms = godd.NewMicroService(app, path, nil, nil)
	ms.Get("/Hello", HandlerHello())

	return ms
}
