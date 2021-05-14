package apilifecycle

// OnStart Set
func (api *APILifeCycle) OnStart(handler HandlerCycle) {
	api.onStart = handler
}

// GetOnStart Get
func (api *APILifeCycle) GetOnStart() HandlerCycle {
	return api.onStart
}
