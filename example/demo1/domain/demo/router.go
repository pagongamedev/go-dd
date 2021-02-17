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
			"en": "example/demo1/domain/demo/i18n/demo.en.toml",
			"th": "example/demo1/domain/demo/i18n/demo.th.toml",
		})

	ms = godd.NewMicroService(app, path, nil, nil, i18n)
	ms.Get("/hello", HandlerHello())

	// repo := godd.EnvironmentSwitcher("localhost", 0, 0, 1, 1, 2, "A", "B", "C", "D")
	// log.Println("repo", repo)

	return ms
}
