package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// ValidateAuth Type
type ValidateAuth = func(context *godd.Context) (roleData interface{}, goddErr *godd.Error)

// ValidateAuth Set
func (api *APILifeCycle) ValidateAuth(handler ValidateAuth) {
	api.validateAuth = handler
}

// GetValidateAuth Get
func (api *APILifeCycle) GetValidateAuth() ValidateAuth {
	return api.validateAuth
}

// Handler Default
func handlerDefaultValidateAuth() ValidateAuth {
	return func(context *godd.Context) (roleData interface{}, goddErr *godd.Error) {
		return nil, nil
	}
}
