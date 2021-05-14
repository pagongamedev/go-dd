package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// HandlerByPass Type
type HandlerByPass = func(context godd.InterfaceContext, service interface{}, serviceOptionList map[string]interface{}) error

// SetHandlerByPassLifeCycle API
func (api *APILifeCycle) SetHandlerByPassLifeCycle(handler HandlerByPass) {
	api.byPass = handler
}
