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
func New(interfaceApp godd.InterfaceApp, path string, context *godd.Context) *MicroService {
	apiMiddleware := &goddAPILifeCycle.APILifeCycle{}

	var http *HTTP

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
	return &MicroService{
		http:          ms.http,
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
