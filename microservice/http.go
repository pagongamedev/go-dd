package microservice

import (
	godd "github.com/pagongamedev/go-dd"
	goddAPI "github.com/pagongamedev/go-dd/api"
	goddAPILifeCycle "github.com/pagongamedev/go-dd/apilifecycle"
)

type HTTP struct {
	http          godd.InterfaceHTTP
	context       *godd.Context
	apiMiddleware *goddAPILifeCycle.APILifeCycle
}

// Add API
func (msHTTP *HTTP) Add(method string, path string, api *goddAPI.API) {
	api.Initial(msHTTP.context, msHTTP.apiMiddleware, goddAPI.HTTPMappingStandardResponse, goddAPI.HTTPMappingStandardError)
	msHTTP.http.Add(method, path, msHTTP.context, api.HandlerLifeCycle())
}

// Get API
func (msHTTP *HTTP) Get(path string, api *goddAPI.API) {
	api.Initial(msHTTP.context, msHTTP.apiMiddleware, goddAPI.HTTPMappingStandardResponse, goddAPI.HTTPMappingStandardError)
	msHTTP.http.Get(path, msHTTP.context, api.HandlerLifeCycle())
}

// Post API
func (msHTTP *HTTP) Post(path string, api *goddAPI.API) {
	api.Initial(msHTTP.context, msHTTP.apiMiddleware, goddAPI.HTTPMappingStandardResponse, goddAPI.HTTPMappingStandardError)
	msHTTP.http.Post(path, msHTTP.context, api.HandlerLifeCycle())
}

// Put API
func (msHTTP *HTTP) Put(path string, api *goddAPI.API) {
	api.Initial(msHTTP.context, msHTTP.apiMiddleware, goddAPI.HTTPMappingStandardResponse, goddAPI.HTTPMappingStandardError)
	msHTTP.http.Put(path, msHTTP.context, api.HandlerLifeCycle())
}

// Patch API
func (msHTTP *HTTP) Patch(path string, api *goddAPI.API) {
	api.Initial(msHTTP.context, msHTTP.apiMiddleware, goddAPI.HTTPMappingStandardResponse, goddAPI.HTTPMappingStandardError)
	msHTTP.http.Patch(path, msHTTP.context, api.HandlerLifeCycle())
}

// Delete API
func (msHTTP *HTTP) Delete(path string, api *goddAPI.API) {
	api.Initial(msHTTP.context, msHTTP.apiMiddleware, goddAPI.HTTPMappingStandardResponse, goddAPI.HTTPMappingStandardError)
	msHTTP.http.Delete(path, msHTTP.context, api.HandlerLifeCycle())
}
