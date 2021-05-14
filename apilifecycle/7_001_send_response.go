package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// SendResponse Type
type SendResponse = func(context godd.InterfaceContext, code int, requestValidated interface{}) (err error)

// SendResponse Set
func (api *APILifeCycle) SendResponse(handler SendResponse) {
	api.sendResponse = handler
}

// GetSendResponse Get
func (api *APILifeCycle) GetSendResponse() SendResponse {
	return api.sendResponse
}

// Handler Default
func handlerDefaultSendResponse() SendResponse {
	return func(context godd.InterfaceContext, code int, responseStandard interface{}) error {
		return context.Response(responseStandard, "", code)
	}
}
