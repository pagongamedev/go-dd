package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// ParseRequest Type
type ParseRequest = func(context *godd.Context) (requestMappingBody interface{}, goddErr *godd.Error)

// ParseRequest Set
func (api *APILifeCycle) ParseRequest(handler ParseRequest) {
	api.parseRequest = handler
}

// GetParseRequest Get
func (api *APILifeCycle) GetParseRequest() ParseRequest {
	return api.parseRequest
}

// Handler Default
func handlerDefaultParseRequest() ParseRequest {
	return func(context *godd.Context) (requestMapping interface{}, goddErr *godd.Error) {
		return nil, nil
	}
}
