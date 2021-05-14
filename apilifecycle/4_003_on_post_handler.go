package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// OnPostHandler Type
type OnPostHandler = func(context godd.InterfaceContext, code int, responseRawIn interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseRawOut interface{}, responsePaginationOut *godd.ResponsePagination, goddErr *godd.Error)

// OnPostHandler func
func (api *APILifeCycle) OnPostHandler(handler OnPostHandler) {
	api.onPostHandler = handler
}

// GetOnPostHandler func
func (api *APILifeCycle) GetOnPostHandler() OnPostHandler {
	return api.onPostHandler
}

// Handler Default
func handlerDefaultOnPostHandler() OnPostHandler {
	return func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePaginationIn *godd.ResponsePagination) (codeOut int, responseRawOut interface{}, responsePaginationOut *godd.ResponsePagination, goddErr *godd.Error) {
		return code, responseRaw, responsePaginationIn, nil
	}
}
