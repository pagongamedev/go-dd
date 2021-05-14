package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// MappingStandardResponse Type
type MappingStandardResponse = func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseMapping interface{}, goddErr *godd.Error)

// MappingStandardResponse Set
func (api *APILifeCycle) MappingStandardResponse(handler MappingStandardResponse) {
	api.mappingStandardResponse = handler
}

// GetMappingStandardResponse Get
func (api *APILifeCycle) GetMappingStandardResponse() MappingStandardResponse {
	return api.mappingStandardResponse
}

// Handler Default
func handlerDefaultMappingStandardResponse() MappingStandardResponse {
	return func(context godd.InterfaceContext, code int, responseRaw interface{}, responsePagination *godd.ResponsePagination) (codeOut int, responseStandard interface{}, goddErr *godd.Error) {
		// response, goddErr := MappingStandard(code, responseRaw, responsePagination)
		return code, responseRaw, goddErr
	}
}
