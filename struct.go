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

// ==============================================================

// Error Struct
type Error struct {
	Code          int // Please Use http.Status
	Error         error
	ErrorValidate *map[string]ErrorValidate
}

// ErrorValidate struct
type ErrorValidate struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Param   string `json:"param,omitempty"`
}

// ErrorValidateReason struct
type ErrorValidateReason struct {
}

// ==============================================================

// RequestPagination Struct
type RequestPagination struct {
	Page     int `json:"page"      default:"1"  swaggertype:"integer"`
	PageSize int `json:"page_size" default:"10" swaggertype:"integer"`
	// DateStart string `json:"date_start"`
	// DateEnd   string `json:"date_end"`
	// SortType  string `json:"sort_type"`
}

// RequestFilter Struct
type RequestFilter struct {
}

// ResponseDataList for Send Response Message to Encode Response
type ResponseDataList struct {
	Success            bool                `json:"success"              swaggertype:"boolean"`
	Message            string              `json:"message"              swaggertype:"string"`
	Data               interface{}         `json:"data,omitempty"       swaggertype:"object"`
	ResponsePagination *ResponsePagination `json:"pagination,omitempty" swaggertype:"object"`
	ResponseError      *ResponseError      `json:"error,omitempty"      swaggertype:"object"` // errors don't define JSON marshaling`
}

// ResponsePagination Struct
type ResponsePagination struct {
	Page       int `json:"page"         swaggertype:"integer"`
	PageSize   int `json:"page_size"    swaggertype:"integer"`
	PageCount  int `json:"page_count"   swaggertype:"integer"`
	ItemCount  int `json:"item_count"   swaggertype:"integer"`
	TotalCount int `json:"total_count"  swaggertype:"integer"`
}

// ResponseError Struct
type ResponseError struct {
	Message  string                    `json:"message" swaggertype:"string"`
	Validate *map[string]ErrorValidate `json:"validate,omitempty" swaggertype:"object"`
}

// ==============================================================
