package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// MappingStandardError Type
type MappingStandardError = func(goddErr *godd.Error) (codeOut int, responseError interface{}, goddErrOut *godd.Error)

// MappingStandardError Set
func (api *APILifeCycle) MappingStandardError(handler MappingStandardError) {
	api.mappingStandardError = handler
}

// GetMappingStandardError Get
func (api *APILifeCycle) GetMappingStandardError() MappingStandardError {
	return api.mappingStandardError
}

// Handler Default
func handlerDefaultMappingStandardError() MappingStandardError {
	return func(goddErr *godd.Error) (codeOut int, responseError interface{}, goddErrOut *godd.Error) {
		// response, goddErr := MappingStandard(code, responseRaw, responsePagination)
		return 400, responseError, goddErr
	}
}
