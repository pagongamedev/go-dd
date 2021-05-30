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

// App func
func (app *AppNetHTTP) App() interface{} {
	return app.app
}

// SetApp func
func (app *AppNetHTTP) SetApp(newApp interface{}) {
	app.app = newApp.(*http.ServeMux)
}

// GetFramework func
func (app *AppNetHTTP) GetFramework() godd.FrameWork {
	return app.framework
}

// Listen func
func (app *AppNetHTTP) Listen(port string, extraList ...interface{}) error {
	var handle http.Handler
	handle = app.app

	if len(extraList) > 0 {
		handle = extraList[0].(http.Handler)
	}

	app.server = &http.Server{Addr: ":" + port, Handler: handle}
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
func (app *AppNetHTTP) Get(path string, context *godd.Context, handleList ...godd.Handler) godd.InterfaceHTTP {
	return app.route("get", path, context, handleList...)
}

// Group func
func (app *AppNetHTTP) Group(path string, context *godd.Context, handleList ...godd.Handler) godd.InterfaceHTTP {
	// app.Handle("/payment/v1/", http.StripPrefix("/payment/v1", payment.MakeHandler(paymentService)))
	return app.route("group", path, context, handleList...)
}

func (app *AppNetHTTP) route(routeType string, path string, context *godd.Context, handleList ...godd.Handler) godd.InterfaceHTTP {
	// var h godd.Handler
	// // var router fiber.Router
	// var routeFunc func(ctx *fiber.Ctx) error

	// routeFunc = nil
	// if len(handleList) > 0 {
	// 	h = handleList[0]
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
