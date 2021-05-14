package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// HandlerByPass Type
type HandlerByPass = func(context *godd.Context) error

// SetHandlerByPassLifeCycle API
func (api *APILifeCycle) SetHandlerByPassLifeCycle(handler HandlerByPass) {
	api.byPass = handler
}
