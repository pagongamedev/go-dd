package demo

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
	goddMicroService "github.com/pagongamedev/go-dd/microservice"

	"golang.org/x/text/language"
)

// Router Func
func Router(app *fiber.App, path string) *goddMicroService.MicroService {

	i18n := godd.NewI18N(language.English, "toml", toml.Unmarshal,
		godd.Map{
			"en": "example/demo1/domain/demo/i18n/active.en.toml",
			"th": "example/demo1/domain/demo/i18n/active.th.toml",
		})

	ms := goddMicroService.New(app, path, nil, nil, i18n)
	ms.Get("/hello", HandlerHello())

	msLogin := middlewareLogin(ms)
	msLogin.Get("/hello2", HandlerHello2())

	return ms
}

func middlewareLogin(ms *goddMicroService.MicroService) *goddMicroService.MicroService {
	msLogin := ms.NewOne()
	msLogin.Override().ValidateAuth(func(context godd.InterfaceContext) (roleData interface{}, goddErr *godd.Error) {
		log.Println("Hello")
		return nil, nil
	})
	// msLogin.AppendMiddlewareOnStart(func(context godd.InterfaceContext) (goddErr *godd.Error) {
	// 	log.Println("Start : 1")
	// 	return nil
	// })
	// msLogin.AppendMiddlewareOnStart(func(context godd.InterfaceContext) (goddErr *godd.Error) {
	// 	log.Println("Start : 2")
	// 	return nil
	// })
	// msLogin.AppendMiddlewareOnStart(func(context godd.InterfaceContext) (goddErr *godd.Error) {
	// 	log.Println("Start : 3")
	// 	return nil
	// })

	// msLogin.AppendMiddlewareOnEnd(func(context godd.InterfaceContext) (goddErr *godd.Error) {
	// 	log.Println("End : 1")
	// 	return nil
	// })
	// msLogin.AppendMiddlewareOnEnd(func(context godd.InterfaceContext) (goddErr *godd.Error) {
	// 	log.Println("End : 2")
	// 	return nil
	// })
	// msLogin.AppendMiddlewareOnEnd(func(context godd.InterfaceContext) (goddErr *godd.Error) {
	// 	log.Println("End : 3")
	// 	return nil
	// })

	return msLogin

}
