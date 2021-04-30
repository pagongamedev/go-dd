package godd

import (
	"errors"
	"strings"
)

// Map type
type Map map[string]interface{}

// MapString type
type MapString map[string]string

// DeferClose struct
type DeferClose struct {
	Name string
	I    InterfaceClose
}

type FuncEnvironment = func(secret MapString) (*Map, *[]DeferClose)

// Handler type
type Handler func(InterfaceContext) error

// FrameWork type
type FrameWork string

const (
	// FrameWorkGofiber FrameWork
	FrameWorkGofiber FrameWork = "gofiber"
)

// Env type
type Env string

const (
	// EnvLocalhost Env
	EnvLocalhost Env = "localhost"
	// EnvDevelopment Env
	EnvDevelopment Env = "development"
	// EnvTesting Env
	EnvTesting Env = "testing"
	// EnvStaging Env
	EnvStaging Env = "staging"
	// EnvProduction Env
	EnvProduction Env = "production"
)

// ==============================================================

// Error Struct
type Error struct {
	Code          int // Please Use http.Status
	Error         error
	ErrorValidate *map[string]ResponseErrorValidate
}

func (e *Error) IsContain(subString string) bool {
	return strings.Contains(e.Error.Error(), subString)
}

func (e *Error) SetError(str string) {
	e.Error = errors.New(str)
}

func (e *Error) IsContainSetError(subString string, errorString string) {
	if e.IsContain(subString) {
		e.SetError(errorString)
	}
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
