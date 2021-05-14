package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// OnPreHandler Type
type OnPreHandler = func(context *godd.Context, requestValidatedBody interface{}, requestValidatedParam interface{}, requestValidatedQuery interface{}) (requestValidatedBodyOut interface{}, requestValidatedParamOut interface{}, requestValidatedQueryOut interface{}, goddErr *godd.Error)

// OnPreHandler Set
func (api *APILifeCycle) OnPreHandler(handler OnPreHandler) {
	api.onPreHandler = handler
}

// GetOnPreHandler Get
func (api *APILifeCycle) GetOnPreHandler() OnPreHandler {
	return api.onPreHandler
}

// Handler Default
func handlerDefaultOnPreHandler() OnPreHandler {
	return func(context *godd.Context, requestValidatedBody, requestValidatedParam, requestValidatedQuery interface{}) (requestValidatedBodyOut interface{}, requestValidatedParamOut interface{}, requestValidatedQueryOut interface{}, goddErr *godd.Error) {
		return requestValidatedBody, requestValidatedParam, requestValidatedQuery, nil
	}
}
