package apilifecycle

// AppendMiddlewareEndList func
func (api *APILifeCycle) AppendMiddlewareEndList(handler OnPreResponse) {
	api.middlewareEndList = append(api.middlewareEndList, handler)
}
