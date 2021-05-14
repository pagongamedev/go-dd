package apilifecycle

// AppendMiddlewareStartList func
func (api *APILifeCycle) AppendMiddlewareStartList(handler HandlerCycle) {
	api.middlewareStartList = append(api.middlewareStartList, handler)
}
