package v2

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
)

//==================== Interface App ====================

func NewApp() (godd.InterfaceApp, *fiber.App) {
	app := fiber.New(fiber.Config{ReadTimeout: time.Second * 5})
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
func (app *AppGofiber) Get(path string, handlers ...godd.Handler) godd.InterfaceRouter {
	var h godd.Handler
	var router fiber.Router
	if len(handlers) > 0 {
		h = handlers[0]
		router = app.app.Get(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	} else {
		router = app.app.Group(path)
	}

	return &RouterGofiber{
		router:    &router,
		framework: app.framework,
	}
}

// Group func
func (app *AppGofiber) Group(path string, handlers ...godd.Handler) godd.InterfaceRouter {
	var h godd.Handler
	var router fiber.Router
	if len(handlers) > 0 {
		h = handlers[0]
		router = app.app.Group(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	} else {
		router = app.app.Group(path)
	}

	return &RouterGofiber{
		router:    &router,
		framework: app.framework,
	}
}

//==================== Interface Router ====================

// RouterGofiber struct
type RouterGofiber struct {
	router    *fiber.Router
	framework godd.FrameWork
}

// Add func
func (router *RouterGofiber) Add(method string, path string, handlers ...func(ctx godd.InterfaceContext) error) {
	var h godd.Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Add(method, path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	}
}

// Get func
func (router *RouterGofiber) Get(path string, handlers ...func(ctx godd.InterfaceContext) error) {
	var h godd.Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Get(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	}
}

// Post func
func (router *RouterGofiber) Post(path string, handlers ...func(ctx godd.InterfaceContext) error) {
	var h godd.Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Post(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	}
}

// Put func
func (router *RouterGofiber) Put(path string, handlers ...func(ctx godd.InterfaceContext) error) {
	var h godd.Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Put(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	}
}

// Patch func
func (router *RouterGofiber) Patch(path string, handlers ...func(ctx godd.InterfaceContext) error) {
	var h godd.Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Patch(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	}
}

// Delete func
func (router *RouterGofiber) Delete(path string, handlers ...func(ctx godd.InterfaceContext) error) {
	var h godd.Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Delete(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	}
}

//==================== Interface Context ====================

// AdapterContextGofiber Func
func AdapterContextGofiber(ctx interface{}) godd.InterfaceContext {
	return &ContextGofiber{
		ctx: ctx.(*fiber.Ctx),
	}
}

// ContextGofiber struct
type ContextGofiber struct {
	ctx               *fiber.Ctx
	framework         godd.FrameWork
	Service           interface{}
	State             map[string]interface{}
	ServiceOptionList map[string]interface{}
	i18n              *godd.I18N
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
func (context *ContextGofiber) Response(responseDataList interface{}, responseCode ...int) error {
	if len(responseCode) > 0 {
		context.ctx.Status(responseCode[0])
	}
	return context.ctx.JSON(responseDataList)
}

// Redirect func
func (context *ContextGofiber) Redirect(location string, responseCode ...int) error {
	return context.ctx.Redirect(location, responseCode...)
}

//========

// SetContext func
func (context *ContextGofiber) SetContext(service interface{}, serviceOptionList map[string]interface{}, i18n *godd.I18N, state map[string]interface{}) {
	if state == nil {
		state = map[string]interface{}{}
	}
	context.Service = service
	context.ServiceOptionList = serviceOptionList
	context.State = state
	context.i18n = i18n
}

// GetService func
func (context *ContextGofiber) GetService() interface{} {
	return context.Service
}

// GetServiceOptionList func
func (context *ContextGofiber) GetServiceOptionList(name string) interface{} {
	if context.ServiceOptionList != nil {
		return context.ServiceOptionList[name]
	}
	log.Println("ServiceOptionList is null")
	return nil
}

// GetState func
func (context *ContextGofiber) GetState(name string) interface{} {
	if context.State != nil {
		return context.State[name]
	}
	return nil
}

// SetState func
func (context *ContextGofiber) SetState(name string, value interface{}) {
	if context.State != nil {
		context.State[name] = value
	} else {
		log.Println("State is null")
	}
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

//===========

//SetLang func
func (context *ContextGofiber) SetLang(lang string) {
	if context.i18n != nil {
		context.i18n.SetLang(lang)
	}
}

//GetLang func
func (context *ContextGofiber) GetLang() string {
	return context.i18n.GetLang()
}

// GetI18N func
func (context *ContextGofiber) GetI18N() *godd.I18N {
	return context.i18n
}

// ValidateStruct func
func (context *ContextGofiber) ValidateStruct(i interface{}, iType map[string]interface{}) *godd.Error {
	return godd.ValidateStruct(context.i18n, i, iType)
}

// SetDefaultStruct func
func (context *ContextGofiber) SetDefaultStruct(i interface{}) interface{} {
	return godd.SetDefaultStruct(i)
}
