package godd

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

//==================== Interface App ====================

// AdapterAppGofiber Func
func AdapterAppGofiber(app interface{}, framework FrameWork) InterfaceApp {
	return &AppGofiber{
		app:       app.(*fiber.App),
		framework: framework,
	}
}

// AppGofiber struct
type AppGofiber struct {
	app       *fiber.App
	framework FrameWork
}

// GetFramework func
func (app *AppGofiber) GetFramework() FrameWork {
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
func (app *AppGofiber) Get(path string, handlers ...Handler) InterfaceRouter {
	var h Handler
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
func (app *AppGofiber) Group(path string, handlers ...Handler) InterfaceRouter {
	var h Handler
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
	framework FrameWork
}

// Add func
func (router *RouterGofiber) Add(method string, path string, handlers ...func(ctx InterfaceContext) error) {
	var h Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Add(method, path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	}
}

// Get func
func (router *RouterGofiber) Get(path string, handlers ...func(ctx InterfaceContext) error) {
	var h Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Get(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	}
}

// Post func
func (router *RouterGofiber) Post(path string, handlers ...func(ctx InterfaceContext) error) {
	var h Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Post(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	}
}

// Put func
func (router *RouterGofiber) Put(path string, handlers ...func(ctx InterfaceContext) error) {
	var h Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Put(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	}
}

// Patch func
func (router *RouterGofiber) Patch(path string, handlers ...func(ctx InterfaceContext) error) {
	var h Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Patch(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	}
}

// Delete func
func (router *RouterGofiber) Delete(path string, handlers ...func(ctx InterfaceContext) error) {
	var h Handler
	if len(handlers) > 0 {
		h = handlers[0]
		(*router.router).Delete(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(ctx))
		})
	}
}

//==================== Interface Context ====================

// AdapterContextGofiber Func
func AdapterContextGofiber(ctx interface{}) InterfaceContext {
	return &ContextGofiber{
		ctx: ctx.(*fiber.Ctx),
	}
}

// ContextGofiber struct
type ContextGofiber struct {
	ctx               *fiber.Ctx
	framework         FrameWork
	Service           interface{}
	State             map[string]interface{}
	ServiceOptionList map[string]interface{}
}

// GetFramework func
func (context *ContextGofiber) GetFramework() FrameWork {
	return context.framework
}

// GetFrameworkContext func
func (context *ContextGofiber) GetFrameworkContext() interface{} {
	return context.ctx
}

// Response func
func (context *ContextGofiber) Response(responseDataList interface{}, responseCode ...int) error {
	return context.ctx.JSON(responseDataList)
}

// Redirect func
func (context *ContextGofiber) Redirect(location string, responseCode ...int) error {
	return context.ctx.Redirect(location, responseCode...)
}

//========

// SetContext func
func (context *ContextGofiber) SetContext(api *APIHTTP, state map[string]interface{}) {
	if state == nil {
		state = map[string]interface{}{}
	}
	context.Service = api.service
	context.ServiceOptionList = api.serviceOptionList
	context.State = state
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
