package godd

//==================================

// InterfaceClose for Manage Defer Close
type InterfaceClose interface {
	Close() error
}

//==================================

// InterfaceApp interface
type InterfaceApp interface {
	GetFramework() FrameWork
	GetFrameworkApp() interface{}
	Listen(port string) error
	Shutdown() error
	Get(path string, handlers ...Handler) InterfaceRouter
	Group(path string, handlers ...Handler) InterfaceRouter
}

// InterfaceRouter interface
type InterfaceRouter interface {
	Add(method string, path string, handlers ...func(ctx InterfaceContext) error)
	Get(path string, handlers ...func(ctx InterfaceContext) error)
	Post(path string, handlers ...func(ctx InterfaceContext) error)
	Put(path string, handlers ...func(ctx InterfaceContext) error)
	Patch(path string, handlers ...func(ctx InterfaceContext) error)
	Delete(path string, handlers ...func(ctx InterfaceContext) error)
}

// InterfaceContext interface
type InterfaceContext interface {
	GetFramework() FrameWork
	GetFrameworkContext() interface{}
	Response(responseDataList interface{}, responseCode ...int) error
	Redirect(location string, responseCode ...int) error

	SetContext(api *APIHTTP, state map[string]interface{})
	GetService() interface{}
	GetServiceOptionList(name string) interface{}
	GetState(name string) interface{}
	SetState(name string, value interface{})

	SetContentType(str string)
}

// Get(path string, handlers ...Handler) Router

//==================================
