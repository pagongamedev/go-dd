package demo

import (
	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
	"golang.org/x/text/language"
)

var ms *godd.MicroService

// Router Func
func Router(app *fiber.App, path string) *godd.MicroService {

	i18n := godd.NewI18N(language.English, "toml", toml.Unmarshal,
		godd.Map{
			"en-US": "example/demo1/domain/demo/i18n/demo.en.toml",
			"th-TH": "example/demo1/domain/demo/i18n/demo.th.toml",
		})

	ms = godd.NewMicroService(app, path, nil, nil, i18n)
	ms.Get("/hello", HandlerHello())

	return ms
}
