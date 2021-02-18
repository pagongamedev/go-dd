package demo

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
	goddMicroService "github.com/pagongamedev/go-dd/microservice"

	"golang.org/x/text/language"
)

var ms *goddMicroService.MicroService

// Router Func
func Router(app *fiber.App, path string) *goddMicroService.MicroService {

	i18n := godd.NewI18N(language.English, "toml", toml.Unmarshal,
		godd.Map{
			"en": "example/demo1/domain/demo/i18n/demo.en.toml",
			"th": "example/demo1/domain/demo/i18n/demo.th.toml",
		})

	ms = goddMicroService.New(app, path, nil, nil, i18n)
	// ms.Get("/hello", HandlerHello())
	ms.Get("/hello", HandlerHello())

	msLogin := ms.NewOne()
	msLogin.Override().ValidateAuth(func(context godd.InterfaceContext) (err *godd.Error) {
		log.Println("Hello")
		return nil
	})

	msLogin.Get("/hello2", HandlerHello2())

	//msLogin.Get()
	// repo := godd.EnvironmentSwitcher("localhost", 0, 0, 1, 1, 2, "A", "B", "C", "D")
	// log.Println("repo", repo)

	return ms
}
