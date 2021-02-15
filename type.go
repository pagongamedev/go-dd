package godd

// Map type
type Map map[string]interface{}

// Handler type
type Handler func(InterfaceContext) error

// FrameWork type
type FrameWork string

var (
	// FrameWorkGofiber FrameWork
	FrameWorkGofiber FrameWork = "gofiber"
)
