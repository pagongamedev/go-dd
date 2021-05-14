package godd

import (
	"errors"
	"log"
	"net/http"
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
type Handler func(*Context) error

// FrameWork type
type FrameWork string

const (
	// FrameWorkGofiber FrameWork
	FrameWorkGofiberV2 FrameWork = "gofiber/v2"
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

type Context struct {
	i                 InterfaceContext
	service           interface{}
	serviceOptionList map[string]interface{}
	state             map[string]interface{}
	i18n              *I18N
}

func NewContext(i InterfaceContext, service interface{}, serviceOptionList map[string]interface{}, state map[string]interface{}, i18n *I18N) *Context {

	if serviceOptionList == nil {
		serviceOptionList = map[string]interface{}{}
	}

	if state == nil {
		state = map[string]interface{}{}
	}

	return &Context{i, service, serviceOptionList, state, i18n}
}

func (context *Context) App() InterfaceContext {
	return context.i
}

// GetService func
func (context *Context) GetService() interface{} {
	return context.service
}

// GetServiceOptionList func
func (context *Context) GetServiceOptionList(name string) interface{} {
	if context.serviceOptionList != nil {
		return context.serviceOptionList[name]
	}
	log.Println("ServiceOptionList is null")
	return nil
}

// GetState func
func (context *Context) GetState(name string) interface{} {
	if context.state == nil {
		context.state = map[string]interface{}{}
	}

	return context.state[name]
}

// SetState func
func (context *Context) SetState(name string, value interface{}) {
	if context.state == nil {
		context.state = map[string]interface{}{}
	}

	context.state[name] = value
}

// ClearState func
func (context *Context) ClearState() {
	context.state = map[string]interface{}{}
}

// SetLang func
func (context *Context) SetLang(lang string) {
	if context.i18n != nil {
		context.i18n.SetLang(lang)
	}
}

// GetLang func
func (context *Context) GetLang() string {
	return context.i18n.GetLang()
}

// GetI18N func
func (context *Context) GetI18N() *I18N {
	return context.i18n
}

//===========

// ValidateStruct func
func (context *Context) ValidateStruct(i interface{}, iType map[string]interface{}) *Error {
	return ValidateStruct(context.i18n, i, iType)
}

// SetDefaultStruct func
func (context *Context) SetDefaultStruct(i interface{}) interface{} {
	return SetDefaultStruct(i)
}

// ==============================================================

// Error Struct
type Error struct {
	Code          int // Please Use http.Status
	message       string
	Error         error
	errorValidate *map[string]ResponseErrorValidate
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

func (e *Error) SetMessage(str string) {
	e.message = str
}

func (e *Error) GetMessage() string {
	if e.message != "" {
		return e.message
	}
	if e.errorValidate != nil {
		return http.StatusText(e.Code)
	}

	if e.Error != nil && e.Error.Error() != "" {
		return e.Error.Error()
	} else {
		return http.StatusText(e.Code)
	}
}

func (e *Error) SetErrorValidate(errorValidate *map[string]ResponseErrorValidate) {
	e.Code = http.StatusBadRequest
	e.errorValidate = errorValidate
}

func (e *Error) GetErrorValidate() *map[string]ResponseErrorValidate {
	return e.errorValidate
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
	Code               int                 `json:"code"                 swaggertype:"integer"`
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
