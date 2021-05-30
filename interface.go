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
	SetApp(interface{})
	Listen(port string, extraList ...interface{}) error
	Shutdown() error
	Get(path string, context *Context, handleList ...Handler) InterfaceHTTP
	Group(path string, context *Context, handleList ...Handler) InterfaceHTTP
	IsSupportHTTP() bool
}

//==================================

// InterfaceHTTP interface
type InterfaceHTTP interface {
	Add(method string, path string, context *Context, handleList ...func(context *Context) error)
	Get(path string, context *Context, handleList ...func(context *Context) error)
	Post(path string, context *Context, handleList ...func(context *Context) error)
	Put(path string, context *Context, handleList ...func(context *Context) error)
	Patch(path string, context *Context, handleList ...func(context *Context) error)
	Delete(path string, context *Context, handleList ...func(context *Context) error)
}

//==================================

// InterfaceAMQP interface
type InterfaceAMQP interface {
}

//==================================

// InterfaceContext interface
type InterfaceContext interface {
	GetFramework() FrameWork
	GetFrameworkContext() interface{}
	Response(responseDataList interface{}, contentType string, responseCode ...int) error
	Redirect(location string, responseCode ...int) error

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
}

// Get(path string, handleList ...Handler) Router

//==================================
