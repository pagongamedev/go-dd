package apilifecycle

import godd "github.com/pagongamedev/go-dd"

// ValidateParam Type
type ValidateParam = func(context *godd.Context) (requestValidatedParam interface{}, goddErr *godd.Error)

// ValidateParam Set
func (api *APILifeCycle) ValidateParam(handler ValidateParam) {
	api.validateParam = handler
}

// GetValidateParam Get
func (api *APILifeCycle) GetValidateParam() ValidateParam {
	return api.validateParam
}

// Handler Default
func handlerDefaultValidateParam() ValidateParam {
	return func(context *godd.Context) (requestValidatedParam interface{}, goddErr *godd.Error) {
		return nil, nil
	}
}
