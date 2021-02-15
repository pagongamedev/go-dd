package demo

import (
	godd "github.com/pagongamedev/go-dd"
)

// HandlerHello API
func HandlerHello() *godd.APIHTTP {
	api := godd.NewAPIHTTP()

	api.HandlerLogic(func(context godd.InterfaceContext, requestValidated interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, err *godd.Error) {
		return 200, godd.Map{"message": "Helllllo"}, nil, nil
	})
	return api
}
