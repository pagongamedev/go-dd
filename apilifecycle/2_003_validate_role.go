package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// ValidateRole Type
type ValidateRole = func(context godd.InterfaceContext, roleData interface{}) (goddErr *godd.Error)

// ValidateRole Set
func (api *APILifeCycle) ValidateRole(handler ValidateRole) {
	api.validateRole = handler
}

// GetValidateRole Get
func (api *APILifeCycle) GetValidateRole() ValidateRole {
	return api.validateRole
}

// Handler Default
func handlerDefaultValidateRole() ValidateRole {
	return func(context godd.InterfaceContext, roleData interface{}) (goddErr *godd.Error) {
		return nil
	}
}
