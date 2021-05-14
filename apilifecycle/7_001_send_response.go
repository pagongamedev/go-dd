package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// SendResponse Type
type SendResponse = func(context godd.InterfaceContext, code int, requestValidated interface{}) (goddErr *godd.Error)

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
	return func(context godd.InterfaceContext, code int, responseStandard interface{}) (goddErr *godd.Error) {
		if !godd.IsInterfaceIsNil(responseStandard) {
			if context != nil {
				context.SetContentType("application/json; charset=utf-8")
				context.Response(responseStandard, code)
			}
		}
		return nil
	}
}
