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

func NewWithInitial(context *godd.Context, apiMiddleware *goddAPILifeCycle.APILifeCycle, mappingStandardResponse goddAPILifeCycle.MappingStandardResponse, mappingStandardError goddAPILifeCycle.MappingStandardError) *API {
	api := &API{LifeCycle: &goddAPILifeCycle.APILifeCycle{}}
	api.Initial(context, apiMiddleware, mappingStandardResponse, mappingStandardError)
	return api
}

// func NewWithContext(service interface{}, serviceOptionList map[string]interface{}, state map[string]interface{}, i18n *godd.I18N) *API {
// 	return &API{
// 		LifeCycle: &goddAPILifeCycle.APILifeCycle{},
// 		context:   godd.NewContext(nil, service, serviceOptionList, state, i18n),
// 	}
// }

// Initial API
func (api *API) Initial(context *godd.Context, apiMiddleware *goddAPILifeCycle.APILifeCycle, mappingStandardResponse goddAPILifeCycle.MappingStandardResponse, mappingStandardError goddAPILifeCycle.MappingStandardError) {
	api.context = context
	if apiMiddleware == nil {
		apiMiddleware = &goddAPILifeCycle.APILifeCycle{}
	}
	api.apiMiddleware = apiMiddleware
	api.LifeCycle.CheckerLifeCycle(apiMiddleware, mappingStandardResponse, mappingStandardError)
}

// ================================================================

// HandlerLifeCycle func
func (api *API) HandlerLifeCycle() godd.Handler {
	return func(context *godd.Context) error {
		apiResponse, err := api.LifeCycle.HandlerLifeCycle(context)
		if err != nil {
			return err
		}

		return api.LifeCycle.GetSendResponse()(context, apiResponse.Code, apiResponse.Response)
	}
}

func (api *API) RunLifeCycle() (*goddAPILifeCycle.APIResponse, error) {
	return api.LifeCycle.HandlerLifeCycle(api.context)
}
