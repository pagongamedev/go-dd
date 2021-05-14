package api

import (
	godd "github.com/pagongamedev/go-dd"
	goddAPILifeCycle "github.com/pagongamedev/go-dd/apilifecycle"
)

// API struct
type API struct {
	LifeCycle     *goddAPILifeCycle.APILifeCycle
	apiMiddleware *goddAPILifeCycle.APILifeCycle
	context       *godd.Context
}

func New() *API {
	return &API{LifeCycle: &goddAPILifeCycle.APILifeCycle{}}
}

// SetupHandler API
func (api *API) SetupHandler(context *godd.Context, apiMiddleware *goddAPILifeCycle.APILifeCycle, mappingStandardResponse goddAPILifeCycle.MappingStandardResponse, mappingStandardError goddAPILifeCycle.MappingStandardError) {
	api.context = context
	api.apiMiddleware = apiMiddleware
	api.LifeCycle.CheckerLifeCycle(apiMiddleware, mappingStandardResponse, mappingStandardError)
}

// ================================================================

// HandlerLifeCycle func
func (api *API) HandlerLifeCycle() godd.Handler {
	return func(context godd.InterfaceContext) error {
		apiResponse, err := api.LifeCycle.HandlerLifeCycle(context, api.context)
		if err != nil {
			return err
		}

		return api.LifeCycle.GetSendResponse()(context, apiResponse.Code, apiResponse.Response)
	}
}