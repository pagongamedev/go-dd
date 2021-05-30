package http

import (
	"context"
	"net/http"
	"time"

	godd "github.com/pagongamedev/go-dd"
)

func NewApp() (godd.InterfaceApp, *http.ServeMux) {

	app := http.NewServeMux()
	goddApp := &AppNetHTTP{
		app:       app,
		framework: godd.FrameWorkNetHTTP,
	}
	return goddApp, app
}

// AppNetHTTP struct
type AppNetHTTP struct {
	app       *http.ServeMux
	server    *http.Server
	framework godd.FrameWork
}

// GetFramework func
func (app *AppNetHTTP) App() interface{} {
	return app.app
}

// GetFramework func
func (app *AppNetHTTP) GetFramework() godd.FrameWork {
	return app.framework
}

// Listen func
func (app *AppNetHTTP) Listen(port string) error {
	app.server = &http.Server{Addr: ":" + port, Handler: app.app}
	return app.server.ListenAndServe()
}

// Shutdown func
func (app *AppNetHTTP) Shutdown() error {
	// return app.server.Close()
	d := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	return app.server.Shutdown(ctx)
}

// Get func
func (app *AppNetHTTP) Get(path string, context *godd.Context, handlers ...godd.Handler) godd.InterfaceHTTP {
	return app.route("get", path, context, handlers...)
}

// Group func
func (app *AppNetHTTP) Group(path string, context *godd.Context, handlers ...godd.Handler) godd.InterfaceHTTP {
	return app.route("group", path, context, handlers...)
}

func (app *AppNetHTTP) route(routeType string, path string, context *godd.Context, handlers ...godd.Handler) godd.InterfaceHTTP {
	// var h godd.Handler
	// // var router fiber.Router
	// var routeFunc func(ctx *fiber.Ctx) error

	// routeFunc = nil
	// if len(handlers) > 0 {
	// 	h = handlers[0]
	// 	Handler
	// 	routeFunc = func(ctx *fiber.Ctx) error {
	// 		return h(AdapterContextGofiber(context, ctx))
	// 	}
	// }

	// switch routeType {
	// case "get":
	// 	// router = app.app.Handle(path, routeFunc)
	// 	break
	// 	// case "group":
	// 	// 	router = app.app.Group(path, routeFunc)
	// 	// 	break
	// }

	// return &RouterGofiber{
	// 	router:    &router,
	// 	framework: app.framework,
	// }

	return nil
}

// =============================================================================
func (app *AppNetHTTP) IsSupportHTTP() bool {
	return false
}