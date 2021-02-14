package godd

import (
	"github.com/gofiber/fiber/v2"
)

// MicroService Struct
type MicroService struct {
	router            fiber.Router
	service           interface{}
	serviceOptionList map[string]interface{}
}

// NewMicroService API
func NewMicroService(app *fiber.App, path string, service interface{}, serviceOptionList map[string]interface{}) *MicroService {
	router := app.Group(path)
	return &MicroService{
		router:            router,
		service:           service,
		serviceOptionList: serviceOptionList,
	}
}

// Add API
func (ms *MicroService) Add(method string, path string, api *APIHTTP) {
	ms.router.Add(method, path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList))
}

// Get API
func (ms *MicroService) Get(path string, api *APIHTTP) {
	ms.router.Get(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList))
}

// Post API
func (ms *MicroService) Post(path string, api *APIHTTP) {
	ms.router.Post(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList))
}

// Put API
func (ms *MicroService) Put(path string, api *APIHTTP) {
	ms.router.Put(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList))
}

// Patch API
func (ms *MicroService) Patch(path string, api *APIHTTP) {
	ms.router.Patch(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList))
}

// Delete API
func (ms *MicroService) Delete(path string, api *APIHTTP) {
	ms.router.Delete(path, api.SetupHandlerHTTP(ms.service, ms.serviceOptionList))
}
