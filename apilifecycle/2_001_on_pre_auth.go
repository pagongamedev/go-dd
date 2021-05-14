package apilifecycle

// OnPreAuth Set
func (api *APILifeCycle) OnPreAuth(handler HandlerCycle) {
	api.onPreAuth = handler
}

// GetOnPreAuth Get
func (api *APILifeCycle) GetOnPreAuth() HandlerCycle {
	return api.onPreAuth
}
