package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// HandlerLogic Type
type HandlerLogic = func(context *godd.Context, requestValidatedBody interface{}, requestValidatedParam interface{}, requestValidatedQuery interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, goddErr *godd.Error)

// HandlerLogic Set
func (api *APILifeCycle) HandlerLogic(handler HandlerLogic) {
	api.handlerLogic = handler
}

// GetHandlerLogic Get
func (api *APILifeCycle) GetHandlerLogic() HandlerLogic {
	return api.handlerLogic
}

// Handler Default
func handlerDefaultHandlerLogic() HandlerLogic {
	return func(context *godd.Context, requestValidatedBody, requestValidatedParam, requestValidatedQuery interface{}) (code int, responseRaw interface{}, responsePagination *godd.ResponsePagination, goddErr *godd.Error) {
		return 200, nil, nil, nil
	}
}
