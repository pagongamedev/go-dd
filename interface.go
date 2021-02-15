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
	Get(path string, handler Handler)
}

// *fiber.Ctx

// InterfaceContext interface
type InterfaceContext interface {
	GetFramework() FrameWork
	GetFrameworkContext() interface{}
	Response(responseDataList Map, responseCode ...int) error
	Redirect(location string, responseCode ...int) error
}

// Get(path string, handlers ...Handler) Router

//==================================
