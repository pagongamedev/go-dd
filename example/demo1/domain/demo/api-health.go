package demo

import godd "github.com/pagongamedev/go-dd"

// HandlerHealth API
func HandlerHealth() *godd.APIHTTP {
	api := godd.NewAPIHTTP()

	api.HandlerLogic(func(context *godd.Context, requestValidated interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, err *godd.Error) {
		context.Ctx.SendString("Helllllo")
		return 200, nil, nil, nil
	})

	return api
}
