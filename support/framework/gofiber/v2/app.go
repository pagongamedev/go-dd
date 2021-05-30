package v2

import (
	"time"

	"github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
)

//==================== Interface App ====================

func NewApp(cfg ...fiber.Config) (godd.InterfaceApp, *fiber.App) {
	var config fiber.Config

	if len(cfg) > 0 {
		config = cfg[0]
	} else {
		config = fiber.Config{ReadTimeout: time.Second * 5}
	}

	app := fiber.New(config)
	goddApp := &AppGofiber{
		app:       app,
		framework: godd.FrameWorkGofiberV2,
	}
	return goddApp, app
}

// AppGofiber struct
type AppGofiber struct {
	app       *fiber.App
	framework godd.FrameWork
}

// App func
func (app *AppGofiber) App() interface{} {
	return &app.app
}

// SetApp func
func (app *AppGofiber) SetApp(newApp interface{}) {
	app.app = newApp.(*fiber.App)
}

// GetFramework func
func (app *AppGofiber) GetFramework() godd.FrameWork {
	return app.framework
}

// Listen func
func (app *AppGofiber) Listen(port string) error {
	return app.app.Listen(port)
}

// Shutdown func
func (app *AppGofiber) Shutdown() error {
	return app.app.Shutdown()
}

// Get func
func (app *AppGofiber) Get(path string, context *godd.Context, handlers ...godd.Handler) godd.InterfaceHTTP {
	return app.route("get", path, context, handlers...)
}

// Group func
func (app *AppGofiber) Group(path string, context *godd.Context, handlers ...godd.Handler) godd.InterfaceHTTP {
	return app.route("group", path, context, handlers...)
}

func (app *AppGofiber) route(routeType string, path string, context *godd.Context, handlers ...godd.Handler) godd.InterfaceHTTP {
	var h godd.Handler
	var router fiber.Router
	var routeFunc func(ctx *fiber.Ctx) error

	routeFunc = nil
	if len(handlers) > 0 {
		h = handlers[0]
		routeFunc = func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(context, ctx))
		}
	}

	switch routeType {
	case "get":
		router = app.app.Get(path, routeFunc)
		break
	case "group":
		router = app.app.Group(path, routeFunc)
		break
	}

	return &RouterGofiber{
		router:    &router,
		framework: app.framework,
	}
}

// =============================================================================
func (app *AppGofiber) IsSupportHTTP() bool {
	return true
}
