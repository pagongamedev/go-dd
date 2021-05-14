package apilifecycle

// OnPreEndList Set
func (api *APILifeCycle) OnPreEndList(handler OnPreResponse) {
	api.onPreEndList = handler
}

// GetOnPreEndList Get
func (api *APILifeCycle) GetOnPreEndList() OnPreResponse {
	return api.onPreEndList
}
