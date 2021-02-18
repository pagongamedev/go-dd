package microservice

import (
	godd "github.com/pagongamedev/go-dd"
	"github.com/pagongamedev/go-dd/api"
	"github.com/pagongamedev/go-dd/framework"
	mdw "github.com/pagongamedev/go-dd/middleware"
)

// MicroService Struct
type MicroService struct {
	router            godd.InterfaceRouter
	service           interface{}
	serviceOptionList map[string]interface{}
	i18n              *godd.I18N
	middleware        *mdw.Middleware
}

// New API
func New(app interface{}, path string, service interface{}, serviceOptionList map[string]interface{}, i18n *godd.I18N, fw ...godd.FrameWork) *MicroService {

	interfaceApp := framework.AdapterApp(app, fw...)
	router := interfaceApp.Group(path)
	return &MicroService{
		router:            router,
		service:           service,
		serviceOptionList: serviceOptionList,
		i18n:              i18n,
		middleware:        &mdw.Middleware{LifeCycle: &godd.APILifeCycle{}},
	}
}

// Add API
func (ms *MicroService) Add(method string, path string, api *api.HTTP) {
	api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n, ms.middleware)
	ms.router.Add(method, path, api.HandlerLifeCycle())
}

// Get API
func (ms *MicroService) Get(path string, api *api.HTTP) {
	api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n, ms.middleware)
	ms.router.Get(path, api.HandlerLifeCycle())
}

// Post API
func (ms *MicroService) Post(path string, api *api.HTTP) {
	api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n, ms.middleware)
	ms.router.Post(path, api.HandlerLifeCycle())
}

// Put API
func (ms *MicroService) Put(path string, api *api.HTTP) {
	api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n, ms.middleware)
	ms.router.Put(path, api.HandlerLifeCycle())
}

// Patch API
func (ms *MicroService) Patch(path string, api *api.HTTP) {
	api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n, ms.middleware)
	ms.router.Patch(path, api.HandlerLifeCycle())
}

// Delete API
func (ms *MicroService) Delete(path string, api *api.HTTP) {
	api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n, ms.middleware)
	ms.router.Delete(path, api.HandlerLifeCycle())
}

// NewOne is New Microservice with Clear Middleware
func (ms *MicroService) NewOne() *MicroService {
	return &MicroService{
		router:            ms.router,
		service:           ms.service,
		serviceOptionList: ms.serviceOptionList,
		i18n:              ms.i18n,
		middleware:        &mdw.Middleware{LifeCycle: &godd.APILifeCycle{}},
	}
}

// Override Func API
func (ms *MicroService) Override() *godd.APILifeCycle {
	return ms.middleware.LifeCycle
}

// AppendMiddlewareOnStart func
func (ms *MicroService) AppendMiddlewareOnStart(handler godd.HandlerCycle) {
	ms.middleware.HandlerStartList = append(ms.middleware.HandlerStartList, handler)
}

// AppendMiddlewareOnEnd func
func (ms *MicroService) AppendMiddlewareOnEnd(handler godd.HandlerCycle) {
	ms.middleware.HandlerEndList = append(ms.middleware.HandlerEndList, handler)
}
