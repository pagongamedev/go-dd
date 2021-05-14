package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// ValidateQuery Type
type ValidateQuery = func(context godd.InterfaceContext) (requestValidatedQuery interface{}, goddErr *godd.Error)

// ValidateQuery Set
func (api *APILifeCycle) ValidateQuery(handler ValidateQuery) {
	api.validateQuery = handler
}

// GetValidateQuery Get
func (api *APILifeCycle) GetValidateQuery() ValidateQuery {
	return api.validateQuery
}

// Handler Default
func handlerDefaultValidateQuery() ValidateQuery {
	return func(context godd.InterfaceContext) (requestValidatedQuery interface{}, goddErr *godd.Error) {
		return nil, nil
	}
}
