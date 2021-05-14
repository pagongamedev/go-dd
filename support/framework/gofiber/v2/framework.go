package v2

import (
	"log"
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

// GetFramework func
func (app *AppGofiber) App() interface{} {
	return app.app
}

// GetFramework func
func (app *AppGofiber) GetFramework() godd.FrameWork {
	return app.framework
}

// GetFrameworkApp func
func (app *AppGofiber) GetFrameworkApp() interface{} {
	return &app.app
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
	var h godd.Handler
	var router fiber.Router
	if len(handlers) > 0 {
		h = handlers[0]
		router = app.app.Get(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(context, ctx))
		})
	} else {
		router = app.app.Get(path)
	}

	return &RouterGofiber{
		router:    &router,
		framework: app.framework,
	}
}

// Group func
func (app *AppGofiber) Group(path string, context *godd.Context, handlers ...godd.Handler) godd.InterfaceHTTP {
	var h godd.Handler
	var router fiber.Router
	if len(handlers) > 0 {
		h = handlers[0]
		router = app.app.Group(path, func(ctx *fiber.Ctx) error {

			return h(AdapterContextGofiber(context, ctx))
		})
	} else {
		router = app.app.Group(path)
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

// =============================================================================

//==================== Interface Router ====================

// RouterGofiber struct
type RouterGofiber struct {
	router    *fiber.Router
	framework godd.FrameWork
}

// Add func
func (router *RouterGofiber) Add(method string, path string, context *godd.Context, handlers ...func(context *godd.Context) error) {
	var h godd.Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Add(method, path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(context, ctx))
		})
	}
}

// Get func
func (router *RouterGofiber) Get(path string, context *godd.Context, handlers ...func(context *godd.Context) error) {
	var h godd.Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Get(path, func(ctx *fiber.Ctx) error {

			return h(AdapterContextGofiber(context, ctx))
		})
	}
}

// Post func
func (router *RouterGofiber) Post(path string, context *godd.Context, handlers ...func(context *godd.Context) error) {
	var h godd.Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Post(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(context, ctx))
		})
	}
}

// Put func
func (router *RouterGofiber) Put(path string, context *godd.Context, handlers ...func(context *godd.Context) error) {
	var h godd.Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Put(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(context, ctx))
		})
	}
}

// Patch func
func (router *RouterGofiber) Patch(path string, context *godd.Context, handlers ...func(context *godd.Context) error) {
	var h godd.Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Patch(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(context, ctx))
		})
	}
}

// Delete func
func (router *RouterGofiber) Delete(path string, context *godd.Context, handlers ...func(context *godd.Context) error) {
	var h godd.Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Delete(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(context, ctx))
		})
	}
}

//==================== Interface Context ====================

// AdapterContextGofiber Func
func AdapterContextGofiber(context *godd.Context, ctx interface{}) *godd.Context {
	if context == nil {
		context = godd.NewContextDefault()
	}

	context.SetApp(&ContextGofiber{
		ctx: ctx.(*fiber.Ctx),
	})

	return context
}

// ContextGofiber struct
type ContextGofiber struct {
	ctx       *fiber.Ctx
	framework godd.FrameWork
}

// GetFramework func
func (context *ContextGofiber) GetFramework() godd.FrameWork {
	return context.framework
}

// GetFrameworkContext func
func (context *ContextGofiber) GetFrameworkContext() interface{} {
	return context.ctx
}

// Response func
func (context *ContextGofiber) Response(responseDataList interface{}, contentType string, responseCode ...int) error {
	if contentType == "" {
		context.ctx.Context().SetContentType(contentType)
	}
	if len(responseCode) > 0 {
		context.ctx.Status(responseCode[0])
	}
	return context.ctx.JSON(responseDataList)
}

// Redirect func
func (context *ContextGofiber) Redirect(location string, responseCode ...int) error {
	return context.ctx.Redirect(location, responseCode...)
}

//===========

//SetContentType func
func (context *ContextGofiber) SetContentType(str string) {
	context.ctx.Context().SetContentType(str)
}

//SetHeader func
func (context *ContextGofiber) SetHeader(key string, val string) {
	context.ctx.Set(key, val)
}

//GetHeader func
func (context *ContextGofiber) GetHeader(key string, defaultValue ...string) string {
	return context.ctx.Get(key, defaultValue...)
}

//GetQuery func
func (context *ContextGofiber) GetQuery(key string, defaultValue ...string) string {
	return context.ctx.Query(key, defaultValue...)
}

//QueryParser func
func (context *ContextGofiber) QueryParser(out interface{}) error {
	return context.ctx.QueryParser(out)
}

//GetParam func
func (context *ContextGofiber) GetParam(key string, defaultValue ...string) string {
	return context.ctx.Params(key, defaultValue...)
}

//GetBody func
func (context *ContextGofiber) GetBody() []byte {
	return context.ctx.Body()
}

//BodyParser func
func (context *ContextGofiber) BodyParser(out interface{}) error {
	return context.ctx.BodyParser(out)
}

//GetCookie func
func (context *ContextGofiber) GetCookie(key string, val string) {
	context.ctx.Cookies(key, val)
}

//SetCookie func
func (context *ContextGofiber) SetCookie(cookie interface{}) {
	c := cookie.(*fiber.Cookie)
	context.ctx.Cookie(c)
}

//ClearCookie func
func (context *ContextGofiber) ClearCookie(key ...string) {
	context.ctx.ClearCookie(key...)
}

//===========

// Log func
func (context *ContextGofiber) Log(v ...interface{}) {
	log.Println(v...)
}
