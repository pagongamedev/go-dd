package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// OnPreResponse Type
type OnPreResponse = func(context *godd.Context, code int, requestValidatedIn interface{}) (codeOut int, requestValidatedOut interface{}, goddErr *godd.Error)

// OnPreResponse Set
func (api *APILifeCycle) OnPreResponse(handler OnPreResponse) {
	api.onPreResponse = handler
}

// GetOnPreResponse func
func (api *APILifeCycle) GetOnPreResponse() OnPreResponse {
	return api.onPreResponse
}

// Handler Default
func handlerDefaultOnPreResponse() OnPreResponse {
	return func(context *godd.Context, code int, responseStandard interface{}) (codeOut int, responseStandardOut interface{}, goddErr *godd.Error) {
		return code, responseStandard, nil
	}
}
