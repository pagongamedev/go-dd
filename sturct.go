package godd

import "github.com/gofiber/fiber/v2"

// Context Struct
type Context struct {
	Ctx               *fiber.Ctx
	Service           interface{}
	State             map[string]interface{}
	ServiceOptionList map[string]interface{}
}

// Error Struct
type Error struct {
	Code  int // Please Use http.Status
	Error error
}

// ==============================================================

// RequestPagination Struct
type RequestPagination struct {
	Page     int `json:"page"      swaggertype:"integer"`
	PageSize int `json:"page_size" swaggertype:"integer"`
	// DateStart string `json:"date_start"`
	// DateEnd   string `json:"date_end"`
	// SortType  string `json:"sort_type"`
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
	Message string `json:"message" swaggertype:"string"`
}
