package godd

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// APIHTTP Struct
type APIHTTP struct {
	handlerByPass     HandlerByPass
	api               *apiLifeCycle
	service           interface{}
	serviceOptionList map[string]interface{}
}

// apiLifeCycle Struct
type apiLifeCycle struct {
	OnPreAuth               Handler
	ValidateAuth            Handler
	OnPostAuth              Handler
	ValidateHeaders         Handler
	ValidateParams          Handler
	ValidateQuery           Handler
	ParseRequest            ParseRequest
	ValidateRequest         ValidateRequest
	OnPreHandler            OnPreHandler
	HandlerLogic            HandlerLogic
	OnPostHandler           OnPostHandler
	MappingResponse         MappingResponse
	ValidateResponse        ValidateResponse
	MappingResponseStandard MappingResponseStandard
	OnPreResponse           OnPreResponse
	SendResponse            SendResponse
}

// NewAPI API
func NewAPIHTTP() *APIHTTP {
	api := apiLifeCycle{
		OnPreAuth:               handlerDefault(),
		ValidateAuth:            handlerDefault(),
		OnPostAuth:              handlerDefault(),
		ValidateHeaders:         handlerDefault(),
		ValidateParams:          handlerDefault(),
		ValidateQuery:           handlerDefault(),
		ParseRequest:            handlerParseRequestDefault(),
		ValidateRequest:         handlerValidateRequestDefault(),
		OnPreHandler:            handlerOnPreHandlerDefault(),
		HandlerLogic:            handlerHandlerLogicDefault(),
		OnPostHandler:           handlerOnPostHandlerDefault(),
		MappingResponse:         handlerMappingResponseDefault(),
		ValidateResponse:        handlerValidateResponseDefault(),
		MappingResponseStandard: handlerMappingResponseStandardDefault(),
		OnPreResponse:           handlerOnPreResponseDefault(),
		SendResponse:            handlerSendResponseDefault(),
	}

	return &APIHTTP{api: &api}
}

// HandlerByPass Type
type HandlerByPass = func(service interface{}, serviceOptionList map[string]interface{}) fiber.Handler

// Handler Type
type Handler = func(context *Context) (err *Error)

// ParseRequest Type
type ParseRequest = func(context *Context) (requestMapping interface{}, err *Error)

// ValidateRequest Type
type ValidateRequest = func(context *Context, requestMapping interface{}) (requestValidated interface{}, err *Error)

// OnPreHandler Type
type OnPreHandler = func(context *Context, requestValidatedIn interface{}) (requestValidatedOut interface{}, err *Error)

// HandlerLogic Type
type HandlerLogic = func(context *Context, requestValidated interface{}) (code int, responseRaw interface{}, responsePagination *ResponsePagination, err *Error)

// OnPostHandler Type
type OnPostHandler = func(context *Context, code int, responseRawIn interface{}, responsePagination *ResponsePagination) (codeOut int, responseRawOut interface{}, responsePaginationOut *ResponsePagination, err *Error)

// MappingResponse Type
type MappingResponse = func(context *Context, code int, responseRaw interface{}, responsePagination *ResponsePagination) (codeOut int, responseMapping interface{}, responsePaginationOut *ResponsePagination, err *Error)

// ValidateResponse Type
type ValidateResponse = func(context *Context, code int, responseMapping interface{}, responsePagination *ResponsePagination) (codeOut int, responseValidated interface{}, responsePaginationOut *ResponsePagination, err *Error)

// MappingResponseStandard Type
type MappingResponseStandard = func(context *Context, code int, responseRaw interface{}, responsePagination *ResponsePagination) (codeOut int, responseMapping interface{}, err *Error)

// OnPreResponse Type
type OnPreResponse = func(context *Context, code int, requestValidatedIn interface{}) (codeOut int, requestValidatedOut interface{}, err *Error)

// SendResponse Type
type SendResponse = func(context *Context, code int, requestValidated interface{}) (err *Error)

func handlerDefault() Handler {
	return func(context *Context) (err *Error) {
		return nil
	}
}

func handlerParseRequestDefault() ParseRequest {
	return func(context *Context) (requestMapping interface{}, err *Error) {
		return nil, nil
	}
}

func handlerValidateRequestDefault() ValidateRequest {
	return func(context *Context, requestMapping interface{}) (requestValidated interface{}, err *Error) {
		return nil, nil
	}
}

func handlerOnPreHandlerDefault() OnPreHandler {
	return func(context *Context, requestValidatedIn interface{}) (requestValidatedOut interface{}, err *Error) {
		return requestValidatedIn, nil
	}
}

func handlerHandlerLogicDefault() HandlerLogic {
	return func(context *Context, requestValidated interface{}) (code int, responseRaw interface{}, responsePagination *ResponsePagination, err *Error) {
		return 200, nil, nil, nil
	}
}

func handlerOnPostHandlerDefault() OnPostHandler {
	return func(context *Context, code int, responseRawIn interface{}, responsePaginationIn *ResponsePagination) (codeOut int, responseRawOut interface{}, responsePaginationOut *ResponsePagination, err *Error) {
		return code, responseRawIn, responsePaginationIn, nil
	}
}

func handlerMappingResponseDefault() MappingResponse {
	return func(context *Context, code int, responseRaw interface{}, responsePagination *ResponsePagination) (codeOut int, responseMapping interface{}, responsePaginationOut *ResponsePagination, err *Error) {
		return code, responseMapping, responsePagination, nil
	}
}

func handlerValidateResponseDefault() ValidateResponse {
	return func(context *Context, code int, responseMapping interface{}, responsePagination *ResponsePagination) (codeOut int, responseValidated interface{}, responsePaginationOut *ResponsePagination, err *Error) {
		return code, responseMapping, responsePagination, nil
	}
}

func handlerMappingResponseStandardDefault() MappingResponseStandard {
	return func(context *Context, code int, responseRaw interface{}, responsePagination *ResponsePagination) (codeOut int, responseStandard interface{}, err *Error) {

		response, err := MappingStandard(code, responseRaw, responsePagination)
		return code, response, err
	}
}

func handlerOnPreResponseDefault() OnPreResponse {
	return func(context *Context, code int, responseStandard interface{}) (codeOut int, responseStandardOut interface{}, err *Error) {
		return code, responseStandard, nil
	}
}

func handlerSendResponseDefault() SendResponse {
	return func(context *Context, code int, responseStandard interface{}) (err *Error) {
		if responseStandard != nil {
			context.Ctx.Status(code).JSON(responseStandard)
		}
		return nil
	}
}

// ================================================================

// SetHandlerByPassLifeCycle API
func (api *APIHTTP) SetHandlerByPassLifeCycle(handler HandlerByPass) {
	api.handlerByPass = handler
}

// SetupHandlerHTTP API
func (api *APIHTTP) SetupHandlerHTTP(service interface{}, serviceOptionList map[string]interface{}) fiber.Handler {
	api.service = service
	api.service = serviceOptionList
	return api.handlerLifeCycle()
}

// ================================================================

func (api *APIHTTP) handlerLifeCycle() fiber.Handler {
	if api.handlerByPass != nil {
		return api.handlerByPass(api.service, api.serviceOptionList)
	}

	return func(ctx *fiber.Ctx) error {
		// ===============================================

		context := &Context{
			Ctx:               ctx,
			Service:           api.service,
			ServiceOptionList: api.serviceOptionList,
			State:             map[string]interface{}{},
		}

		var err *Error

		// ================== Auth =======================

		err = api.api.OnPreAuth(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.api.ValidateAuth(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.api.OnPostAuth(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		// ================== Validate Request =======================

		err = api.api.ValidateHeaders(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.api.ValidateParams(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.api.ValidateQuery(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		requestMapping, err := api.api.ParseRequest(context)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		requestValidated, err := api.api.ValidateRequest(context, requestMapping)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		// ================== Handler =======================

		requestValidated, err = api.api.OnPreHandler(context, requestValidated)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseRaw, responsePagination, err := api.api.HandlerLogic(context, requestValidated)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseRaw, responsePagination, err = api.api.OnPostHandler(context, code, responseRaw, responsePagination)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		// ================== Validate Response =======================

		code, responseMapping, responsePagination, err := api.api.MappingResponse(context, code, responseRaw, responsePagination)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseValidated, responsePagination, err := api.api.ValidateResponse(context, code, responseMapping, responsePagination)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseStandard, err := api.api.MappingResponseStandard(context, code, responseValidated, responsePagination)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		code, responseStandard, err = api.api.OnPreResponse(context, code, responseStandard)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		err = api.api.SendResponse(context, code, responseStandard)
		if err != nil {
			return encodeErrorHandler(context, err)
		}

		return nil
	}
}

//========================================

// OnPreAuth func
func (api *APIHTTP) OnPreAuth(handler Handler) {
	api.api.OnPreAuth = handler
}

// ValidateAuth func
func (api *APIHTTP) ValidateAuth(handler Handler) {
	api.api.ValidateAuth = handler
}

// ValidateHeaders func
func (api *APIHTTP) ValidateHeaders(handler Handler) {
	api.api.ValidateHeaders = handler
}

// ValidateParams func
func (api *APIHTTP) ValidateParams(handler Handler) {
	api.api.ValidateParams = handler
}

// ValidateQuery func
func (api *APIHTTP) ValidateQuery(handler Handler) {
	api.api.ValidateQuery = handler
}

// ParseRequest func
func (api *APIHTTP) ParseRequest(handler ParseRequest) {
	api.api.ParseRequest = handler
}

// ValidateRequest func
func (api *APIHTTP) ValidateRequest(handler ValidateRequest) {
	api.api.ValidateRequest = handler
}

// OnPreHandler func
func (api *APIHTTP) OnPreHandler(handler OnPreHandler) {
	api.api.OnPreHandler = handler
}

// HandlerLogic func
func (api *APIHTTP) HandlerLogic(handler HandlerLogic) {
	api.api.HandlerLogic = handler
}

// OnPostHandler func
func (api *APIHTTP) OnPostHandler(handler OnPostHandler) {
	api.api.OnPostHandler = handler
}

// MappingResponse func
func (api *APIHTTP) MappingResponse(handler MappingResponse) {
	api.api.MappingResponse = handler
}

// ValidateResponse func
func (api *APIHTTP) ValidateResponse(handler ValidateResponse) {
	api.api.ValidateResponse = handler
}

// OnPreResponse func
func (api *APIHTTP) OnPreResponse(handler OnPreResponse) {
	api.api.OnPreResponse = handler
}

// SendResponse func
func (api *APIHTTP) SendResponse(handler SendResponse) {
	api.api.SendResponse = handler
}

//========================================

func encodeErrorHandler(context *Context, err *Error) error {
	var errorMessage string
	if err.Error != nil {
		errorMessage = err.Error.Error()
	} else {
		errorMessage = http.StatusText(err.Code)
	}
	// errorreporting.LogError(errors.New(errorMessage))
	context.Ctx.Context().SetContentType("application/json; charset=utf-8")

	context.Ctx.Status(err.Code).JSON(ResponseDataList{
		Success: false,
		Message: "unsuccess",
		ResponseError: &ResponseError{
			Message: errorMessage,
		},
	})

	return nil
}

// MappingStandard Func
func MappingStandard(code int, dataList interface{}, responsePagination *ResponsePagination) (response ResponseDataList, err *Error) {
	if dataList == nil {
		dataList = []string{}
	}
	isSuccess := code/100 == 2
	var message string
	if isSuccess {
		message = "success"
	} else {
		message = "unsuccess"
	}
	response = ResponseDataList{
		Success:            isSuccess,
		Message:            message,
		Data:               dataList,
		ResponsePagination: responsePagination,
	}
	return response, nil
}
