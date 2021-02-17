package microservice

import (
	godd "github.com/pagongamedev/go-dd"
	"github.com/pagongamedev/go-dd/api"
	"github.com/pagongamedev/go-dd/framework"
)

// MicroService Struct
type MicroService struct {
	router            godd.InterfaceRouter
	service           interface{}
	serviceOptionList map[string]interface{}
	i18n              *godd.I18N
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
	}
}

// Add API
func (ms *MicroService) Add(method string, path string, api *api.HTTP) {
	ms.router.Add(method, path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n))
}

// Get API
func (ms *MicroService) Get(path string, api *api.HTTP) {
	ms.router.Get(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n))
}

// Post API
func (ms *MicroService) Post(path string, api *api.HTTP) {
	ms.router.Post(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n))
}

// Put API
func (ms *MicroService) Put(path string, api *api.HTTP) {
	ms.router.Put(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n))
}

// Patch API
func (ms *MicroService) Patch(path string, api *api.HTTP) {
	ms.router.Patch(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n))
}

// Delete API
func (ms *MicroService) Delete(path string, api *api.HTTP) {
	ms.router.Delete(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n))
}
