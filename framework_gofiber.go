package godd

import "github.com/gofiber/fiber/v2"

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
func (app *AppGofiber) Get(path string, handler Handler) {
	app.app.Get(path, func(ctx *fiber.Ctx) error {
		return handler(AdapterContextGofiber(ctx))
	})
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
	ctx       *fiber.Ctx
	framework FrameWork
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
func (context *ContextGofiber) Response(responseDataList Map, responseCode ...int) error {
	return context.ctx.JSON(responseDataList)
}

// Redirect func
func (context *ContextGofiber) Redirect(location string, responseCode ...int) error {
	return context.ctx.Redirect(location, responseCode...)
}
