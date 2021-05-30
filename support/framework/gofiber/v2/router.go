package v2

import (
	"github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
)

//==================== Interface Router ====================

// RouterGofiber struct
type RouterGofiber struct {
	router    *fiber.Router
	framework godd.FrameWork
}

// Add func
func (router *RouterGofiber) Add(method string, path string, context *godd.Context, handleList ...func(context *godd.Context) error) {
	var h godd.Handler
	if len(handleList) > 0 {
		h = handleList[0]
		(*router.router).Add(method, path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(context, ctx))
		})
	}
}

// Get func
func (router *RouterGofiber) Get(path string, context *godd.Context, handleList ...func(context *godd.Context) error) {
	var h godd.Handler
	if len(handleList) > 0 {
		h = handleList[0]
		(*router.router).Get(path, func(ctx *fiber.Ctx) error {

			return h(AdapterContextGofiber(context, ctx))
		})
	}
}

// Post func
func (router *RouterGofiber) Post(path string, context *godd.Context, handleList ...func(context *godd.Context) error) {
	var h godd.Handler
	if len(handleList) > 0 {
		h = handleList[0]
		(*router.router).Post(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(context, ctx))
		})
	}
}

// Put func
func (router *RouterGofiber) Put(path string, context *godd.Context, handleList ...func(context *godd.Context) error) {
	var h godd.Handler
	if len(handleList) > 0 {
		h = handleList[0]
		(*router.router).Put(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(context, ctx))
		})
	}
}

// Patch func
func (router *RouterGofiber) Patch(path string, context *godd.Context, handleList ...func(context *godd.Context) error) {
	var h godd.Handler
	if len(handleList) > 0 {
		h = handleList[0]
		(*router.router).Patch(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(context, ctx))
		})
	}
}

// Delete func
func (router *RouterGofiber) Delete(path string, context *godd.Context, handleList ...func(context *godd.Context) error) {
	var h godd.Handler
	if len(handleList) > 0 {
		h = handleList[0]
		(*router.router).Delete(path, func(ctx *fiber.Ctx) error {
			return h(AdapterContextGofiber(context, ctx))
		})
	}
}
