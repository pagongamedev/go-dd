package microservice

import (
	godd "github.com/pagongamedev/go-dd"
	goddAPILifeCycle "github.com/pagongamedev/go-dd/apilifecycle"
)

// MicroService Struct
type MicroService struct {
	apiMiddleware *goddAPILifeCycle.APILifeCycle
	context       *godd.Context
	http          HTTP
}

// New API
func New(interfaceApp godd.InterfaceApp, path string, context *godd.Context) *MicroService {

	http := interfaceApp.Group(path)

	return &MicroService{
		http:          http,
		context:       context,
		apiMiddleware: &goddAPILifeCycle.APILifeCycle{},
	}
}

// NewOne is New Microservice with Clear Middleware
func (ms *MicroService) NewOne() *MicroService {
	return &MicroService{
		http:          ms.http,
		context:       ms.context,
		apiMiddleware: &goddAPILifeCycle.APILifeCycle{},
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

func (ms *MicroService) HTTP() error {
	if ms.context.IsSupportHTTP() {

	}
}
