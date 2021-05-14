package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// ValidateRequest Type
type ValidateRequest = func(context *godd.Context, requestMappingBody interface{}) (requestValidatedBody interface{}, goddErr *godd.Error)

// ValidateRequest Set
func (api *APILifeCycle) ValidateRequest(handler ValidateRequest) {
	api.validateRequest = handler
}

// GetValidateRequest Get
func (api *APILifeCycle) GetValidateRequest() ValidateRequest {
	return api.validateRequest
}

// Handler Default
func handlerDefaultValidateRequest() ValidateRequest {
	return func(context *godd.Context, requestMapping interface{}) (requestValidated interface{}, goddErr *godd.Error) {
		return requestMapping, nil
	}
}
