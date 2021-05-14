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
	App() interface{}
	GetFrameworkApp() interface{}
	Listen(port string) error
	Shutdown() error
	Get(path string, handlers ...Handler) InterfaceHTTP
	Group(path string, handlers ...Handler) InterfaceHTTP
	IsSupportHTTP() bool
}

// InterfaceHTTP interface
type InterfaceHTTP interface {
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
	Response(responseDataList interface{}, contentType string, responseCode ...int) error
	Redirect(location string, responseCode ...int) error

	SetContext(context *Context)
	GetContext() *Context

	SetContentType(str string)

	SetHeader(key string, val string)
	GetHeader(key string, defaultValue ...string) string

	GetQuery(key string, defaultValue ...string) string
	QueryParser(out interface{}) error

	GetParam(key string, defaultValue ...string) string

	GetBody() []byte
	BodyParser(out interface{}) error

	GetCookie(key string, val string)
	SetCookie(cookie interface{})

	ClearCookie(key ...string)

	Log(v ...interface{})

	// ====== Validate Struct

	ValidateStruct(i interface{}, iType map[string]interface{}) *Error
	SetDefaultStruct(i interface{}) interface{}
}

// Get(path string, handlers ...Handler) Router

//==================================
