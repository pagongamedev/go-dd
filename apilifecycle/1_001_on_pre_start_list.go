package apilifecycle

// OnPreStartList Set
func (api *APILifeCycle) OnPreStartList(handler HandlerCycle) {
	api.onPreStartList = handler
}

// GetOnPreStartList Get
func (api *APILifeCycle) GetOnPreStartList() HandlerCycle {
	return api.onPreStartList
}
