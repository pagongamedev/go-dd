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
	ErrorValidate *map[string]ResponseErrorValidate
}

// ==============================================================

// RequestPagination Struct
type RequestPagination struct {
	Page     int `json:"page"      default:"1"  swaggertype:"integer"`
	PageSize int `json:"page_size" default:"10" swaggertype:"integer"`
}

// RequestFilter Struct
type RequestFilter struct {
	DateStart  string              `json:"date_start"`
	DateEnd    string              `json:"date_end"`
	FilterList []RequestFilterType `json:"filter"`
	SortList   []RequestSort       `json:"sort"`
}

// RequestFilterType struct
type RequestFilterType struct {
	Field    string `json:"field"`
	Operator string `json:"op"`
	Value    string `json:"value"`
}

// RequestSort struct
type RequestSort struct {
	Field string `json:"field"`
	By    string `json:"by"`
}

// ==============================================================

// ResponseDataList for Send Response Message to Encode Response
type ResponseDataList struct {
	Success            bool                `json:"success"              swaggertype:"boolean"`
	Message            string              `json:"message"              swaggertype:"string"`
	Data               interface{}         `json:"data,omitempty"       swaggertype:"object"`
	ResponsePagination *ResponsePagination `json:"pagination,omitempty" swaggertype:"object"`
	ResponseError      *ResponseError      `json:"error,omitempty"      swaggertype:"object"`
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
	Message  string                            `json:"message" swaggertype:"string"`
	Validate *map[string]ResponseErrorValidate `json:"validate,omitempty" swaggertype:"object"`
}

// ResponseErrorValidate struct
type ResponseErrorValidate struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
	Param   string `json:"param,omitempty"`
}

// ==============================================================
