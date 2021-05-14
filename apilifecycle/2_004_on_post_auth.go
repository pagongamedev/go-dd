package apilifecycle

// OnPostAuth Set
func (api *APILifeCycle) OnPostAuth(handler HandlerCycle) {
	api.onPostAuth = handler
}

// GetOnPostAuth func
func (api *APILifeCycle) GetOnPostAuth() HandlerCycle {
	return api.onPostAuth
}
