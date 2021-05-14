package microservice

import (
	godd "github.com/pagongamedev/go-dd"
	api "github.com/pagongamedev/go-dd/api"
	apiHTTP "github.com/pagongamedev/go-dd/api/http"
	goddAPILifeCycle "github.com/pagongamedev/go-dd/apilifecycle"
)

type HTTP struct {
	http          godd.InterfaceHTTP
	context       *godd.Context
	apiMiddleware *goddAPILifeCycle.APILifeCycle
}

// Add API
func (msHTTP *HTTP) Add(method string, path string, api *api.API) {
	api.SetupHandler(msHTTP.context, msHTTP.apiMiddleware, apiHTTP.MappingStandardResponse, apiHTTP.MappingStandardError)
	msHTTP.http.Add(method, path, api.HandlerLifeCycle())
}

// Get API
func (msHTTP *HTTP) Get(path string, api *api.API) {
	api.SetupHandler(msHTTP.context, msHTTP.apiMiddleware, apiHTTP.MappingStandardResponse, apiHTTP.MappingStandardError)
	msHTTP.http.Get(path, api.HandlerLifeCycle())
}

// Post API
func (msHTTP *HTTP) Post(path string, api *api.API) {
	api.SetupHandler(msHTTP.context, msHTTP.apiMiddleware, apiHTTP.MappingStandardResponse, apiHTTP.MappingStandardError)
	msHTTP.http.Post(path, api.HandlerLifeCycle())
}

// Put API
func (msHTTP *HTTP) Put(path string, api *api.API) {
	api.SetupHandler(msHTTP.context, msHTTP.apiMiddleware, apiHTTP.MappingStandardResponse, apiHTTP.MappingStandardError)
	msHTTP.http.Put(path, api.HandlerLifeCycle())
}

// Patch API
func (msHTTP *HTTP) Patch(path string, api *api.API) {
	api.SetupHandler(msHTTP.context, msHTTP.apiMiddleware, apiHTTP.MappingStandardResponse, apiHTTP.MappingStandardError)
	msHTTP.http.Patch(path, api.HandlerLifeCycle())
}

// Delete API
func (msHTTP *HTTP) Delete(path string, api *api.API) {
	api.SetupHandler(msHTTP.context, msHTTP.apiMiddleware, apiHTTP.MappingStandardResponse, apiHTTP.MappingStandardError)
	msHTTP.http.Delete(path, api.HandlerLifeCycle())
}
