package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// ValidateResponse Type
type ValidateResponse = func(context godd.InterfaceContext, code int, responseMapping interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseValidated interface{}, responsePaginationOut *godd.ResponsePagination, goddErr *godd.Error)

// ValidateResponse Set
func (api *APILifeCycle) ValidateResponse(handler ValidateResponse) {
	api.validateResponse = handler
}

// GetValidateResponse Get
func (api *APILifeCycle) GetValidateResponse() ValidateResponse {
	return api.validateResponse
}

// Handler Default
func handlerDefaultValidateResponse() ValidateResponse {
	return func(context godd.InterfaceContext, code int, responseMapping interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseValidated interface{}, responsePaginationOut *godd.ResponsePagination, goddErr *godd.Error) {
		return code, responseMapping, responsePagination, nil
	}
}
