package microservice

import (
	godd "github.com/pagongamedev/go-dd"
	goddAPILifeCycle "github.com/pagongamedev/go-dd/apilifecycle"
)

// MicroService Struct
type MicroService struct {
	apiMiddleware *goddAPILifeCycle.APILifeCycle
	context       *godd.Context
	http          *HTTP
}

// New API
func New(interfaceApp godd.InterfaceApp, path string, service interface{}, serviceOptionList map[string]interface{}, i18n *godd.I18N) *MicroService {
	apiMiddleware := &goddAPILifeCycle.APILifeCycle{}
	var http *HTTP

	context := godd.NewContext(nil, service, serviceOptionList, nil, i18n)

	if interfaceApp.IsSupportHTTP() {
		http = &HTTP{
			http:          interfaceApp.Group(path),
			context:       context,
			apiMiddleware: apiMiddleware,
		}
	}

	return &MicroService{
		http:          http,
		context:       context,
		apiMiddleware: apiMiddleware,
	}
}

// NewOne is New Microservice with Clear Middleware
func (ms *MicroService) NewOne() *MicroService {
	apiMiddleware := &goddAPILifeCycle.APILifeCycle{}
	var http *HTTP

	if ms.http != nil {
		http = &HTTP{
			http:          ms.http.http,
			context:       ms.context,
			apiMiddleware: apiMiddleware,
		}
	}

	return &MicroService{
		http:          http,
		context:       ms.context,
		apiMiddleware: apiMiddleware,
	}
}

// Override Func API
func (ms *MicroService) Override() *goddAPILifeCycle.APILifeCycle {
	return ms.apiMiddleware
}

// AppendMiddlewareStartList func
func (ms *MicroService) AppendMiddlewareStartList(handler goddAPILifeCycle.HandlerCycle) {
	ms.apiMiddleware.AppendMiddlewareStartList(handler)
}

// AppendMiddlewareEndList func
func (ms *MicroService) AppendMiddlewareEndList(handler goddAPILifeCycle.OnPreResponse) {
	ms.apiMiddleware.AppendMiddlewareEndList(handler)
}

// ====================================================================================

func (ms *MicroService) HTTP() *HTTP {
	return ms.http
}
