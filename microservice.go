package godd

// MicroService Struct
type MicroService struct {
	router            InterfaceRouter
	service           interface{}
	serviceOptionList map[string]interface{}
	i18n              *I18N
}

// NewMicroService API
func NewMicroService(app interface{}, path string, service interface{}, serviceOptionList map[string]interface{}, i18n *I18N, framework ...FrameWork) *MicroService {

	interfaceApp := AdapterApp(app, framework...)
	router := interfaceApp.Group(path)
	return &MicroService{
		router:            router,
		service:           service,
		serviceOptionList: serviceOptionList,
		i18n:              i18n,
	}
}

// Add API
func (ms *MicroService) Add(method string, path string, api *APIHTTP) {
	ms.router.Add(method, path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n))
}

// Get API
func (ms *MicroService) Get(path string, api *APIHTTP) {
	ms.router.Get(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n))
}

// Post API
func (ms *MicroService) Post(path string, api *APIHTTP) {
	ms.router.Post(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n))
}

// Put API
func (ms *MicroService) Put(path string, api *APIHTTP) {
	ms.router.Put(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n))
}

// Patch API
func (ms *MicroService) Patch(path string, api *APIHTTP) {
	ms.router.Patch(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n))
}

// Delete API
func (ms *MicroService) Delete(path string, api *APIHTTP) {
	ms.router.Delete(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList, ms.i18n))
}
