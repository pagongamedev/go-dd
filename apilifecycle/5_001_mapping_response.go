package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// MappingResponse Type
type MappingResponse = func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseMapping interface{}, responsePaginationOut *godd.ResponsePagination, goddErr *godd.Error)

// MappingResponse Set
func (api *APILifeCycle) MappingResponse(handler MappingResponse) {
	api.mappingResponse = handler
}

// GetMappingResponse Get
func (api *APILifeCycle) GetMappingResponse() MappingResponse {
	return api.mappingResponse
}

// Handler Default
func handlerDefaultMappingResponse() MappingResponse {
	return func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseMapping interface{}, responsePaginationOut *godd.ResponsePagination, goddErr *godd.Error) {
		return code, responseRaw, responsePagination, nil
	}
}
