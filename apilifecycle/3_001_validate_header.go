package apilifecycle

// ValidateHeader Set
func (api *APILifeCycle) ValidateHeader(handler HandlerCycle) {
	api.validateHeader = handler
}

// GetValidateHeader Get
func (api *APILifeCycle) GetValidateHeader() HandlerCycle {
	return api.validateHeader
}
