package v2

import (
	"log"

	"github.com/gofiber/fiber/v2"
	godd "github.com/pagongamedev/go-dd"
)

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
